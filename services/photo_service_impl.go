package services

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type PhotoServiceImpl struct {
	S3Uploader *s3manager.Uploader
}

func (ps *PhotoServiceImpl) UploadPhoto(request *models.PhotoUploadRequest, bucketName string) error {
	_, err := ps.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(request.FileName),
		Body:   request.File,
	})
	return err
}
