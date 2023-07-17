package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SeungyeonHwang/personal-photo-gallery/models"
	"github.com/SeungyeonHwang/personal-photo-gallery/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService services.UserService
	SSMService  services.SSMService
	ClientId    string
}

func (uh *UserHandler) GetClientId() (string, error) {
	if uh.ClientId != "" {
		return uh.ClientId, nil
	}
	clientId, err := uh.SSMService.GetParameter("/cognito/client_id")
	if err != nil {
		return "", err
	}
	uh.ClientId = clientId
	return clientId, nil
}

func (uh *UserHandler) HandleRegistrationRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	regReq := new(models.UserRegistrationRequest)
	err := json.Unmarshal([]byte(request.Body), regReq)
	if err != nil {
		logrus.Errorf("Failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Failed to parse request body: %v", err),
		}, nil
	}

	clientId, err := uh.GetClientId()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to get Cognito client ID: %v", err),
		}, nil
	}

	err = uh.UserService.RegisterUser(regReq, clientId)
	if err != nil {
		logrus.Warnf("Error while registering user %s: %v", regReq.Username, err)
		if awsErr, ok := err.(awserr.Error); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("AWS error while registering user: %v", awsErr.Message()),
			}, nil
		} else {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("Unknown error occurred while registering user: %v", err),
			}, nil
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "User registered successfully",
	}, nil
}

func (uh *UserHandler) HandleConfirmationRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	confirmReq := new(models.UserConfirmationRequest)
	err := json.Unmarshal([]byte(request.Body), confirmReq)
	if err != nil {
		logrus.Errorf("Failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Failed to parse request body: %v", err),
		}, nil
	}

	clientId, err := uh.GetClientId()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to get Cognito client ID: %v", err),
		}, nil
	}

	err = uh.UserService.ConfirmUser(confirmReq.Code, confirmReq.Username, clientId)
	if err != nil {
		logrus.Warnf("Error while confirming user %s: %v", confirmReq.Username, err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Internal Server Error: %v", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "User successfully confirmed",
	}, nil
}

func (uh *UserHandler) HandleLoginRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	loginReq := new(models.UserLoginRequest)
	err := json.Unmarshal([]byte(request.Body), loginReq)
	if err != nil {
		logrus.Errorf("Failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Failed to parse request body: %v", err),
		}, nil
	}

	clientId, err := uh.GetClientId()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to get Cognito client ID: %v", err),
		}, nil
	}

	token, err := uh.UserService.LoginUser(loginReq, clientId)
	if err != nil {
		logrus.Warnf("Error while logging in user %s: %v", loginReq.Username, err)
		if awsErr, ok := err.(awserr.Error); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("AWS error while logging in user: %v", awsErr.Message()),
			}, nil
		} else {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("Unknown error occurred while logging in user: %v", err),
			}, nil
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("User logged in successfully, token: %s", *token),
	}, nil
}

func (uh *UserHandler) HandleLogoutRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logoutReq := new(models.UserGlobalLogoutRequest)
	err := json.Unmarshal([]byte(request.Body), logoutReq)
	if err != nil {
		logrus.Errorf("Failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Failed to parse request body: %v", err),
		}, nil
	}

	err = uh.UserService.LogoutUser(logoutReq)
	if err != nil {
		logrus.Warnf("Error while logging out user: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error while logging out user: %v", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "User successfully logged out",
	}, nil
}
