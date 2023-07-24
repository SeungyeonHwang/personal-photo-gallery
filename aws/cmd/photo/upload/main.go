package main

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/config"
	"github.com/SeungyeonHwang/personal-photo-gallery/handlers"
	"github.com/SeungyeonHwang/personal-photo-gallery/services"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	s3Service := config.NewS3Service()

	photoHandler := &handlers.PhotoHandler{
		PhotoService: &services.PhotoServiceImpl{
			S3Uploader: s3Service.Uploader,
		},
		SSMService: &services.SSMServiceImpl{
			SSMClient: s3Service.SSMClient,
		},
		BucketName: "",
	}

	lambda.Start(photoHandler.HandlePhotoUploadRequest)
}
