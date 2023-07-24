package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/SeungyeonHwang/personal-photo-gallery/models"
	"github.com/SeungyeonHwang/personal-photo-gallery/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

type PhotoHandler struct {
	PhotoService services.PhotoService
	SSMService   services.SSMService
	BucketName   string
}

// TODO: 없으면 unique 한 버킷 만들기
func (ph *PhotoHandler) GetBucketName() (string, error) {
	if ph.BucketName != "" {
		return ph.BucketName, nil
	}

	// TODO: 새로만든 bucket name 지정하기
	bucketName, err := ph.SSMService.GetParameter("/s3/bucket_name")
	if err != nil {
		return "", err
	}
	ph.BucketName = bucketName
	return bucketName, nil
}

func (ph *PhotoHandler) HandlePhotoUploadRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bucketName, err := ph.GetBucketName()
	if err != nil {
		logrus.Errorf("Failed to get S3 bucket name: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to get S3 bucket name: %v", err),
		}, nil
	}
	logrus.Infof("Successfully retrieved bucket name: %s", bucketName)

	// Convert the request body to a byte slice
	requestBody := []byte(request.Body)

	contentType, params, parseErr := mime.ParseMediaType(request.Headers["Content-Type"])
	if parseErr != nil || !strings.HasPrefix(contentType, "multipart/") {
		logrus.Errorf("Failed to parse content type: %v", parseErr)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Expecting a multipart message",
		}, nil
	}

	multipartReader := multipart.NewReader(bytes.NewReader(requestBody), params["boundary"])

	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			// All parts read
			break
		}
		if err != nil {
			logrus.Errorf("Failed to read multipart message: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       fmt.Sprintf("Failed to read multipart message: %v", err),
			}, nil
		}

		// This part is a file
		if part.FileName() != "" {
			// Check if the part's Content-Type is an image
			if !strings.HasPrefix(part.Header.Get("Content-Type"), "image/") {
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusBadRequest,
					Body:       "Invalid file type. Expecting an image file.",
				}, nil
			}

			photoUploadReq := &models.PhotoUploadRequest{
				FileName: part.FileName(),
				File:     part,
			}

			err = ph.PhotoService.UploadPhoto(photoUploadReq, bucketName)
			if err != nil {
				logrus.Errorf("Error while uploading photo: %v", err)
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       fmt.Sprintf("Failed to upload photo: %v", err),
				}, nil
			}
			logrus.Info("Photo uploaded successfully")

			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusCreated,
				Body:       "Photo uploaded successfully",
			}, nil
		}
	}

	logrus.Warn("No file part found in multipart form")
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       "No file part found in multipart form",
	}, nil
}
