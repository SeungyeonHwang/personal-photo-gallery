package main

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/config"
	"github.com/SeungyeonHwang/personal-photo-gallery/handlers"
	"github.com/SeungyeonHwang/personal-photo-gallery/services"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cognitoService := config.NewCognitoService()

	userHandler := &handlers.UserHandler{
		UserService: &services.UserServiceImpl{
			CognitoClient: cognitoService.Client,
		},
		SSMService: &services.SSMServiceImpl{
			SSMClient: cognitoService.SSMClient,
		},
		ClientId: "",
	}

	lambda.Start(userHandler.HandleConfirmationRequest)
}
