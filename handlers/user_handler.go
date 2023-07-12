package handlers

import (
	"net/http"

	"github.com/SeungyeonHwang/personal-photo-gallery/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/labstack/echo/v4"
)

type UserRegistrationRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func RegisterUser(c echo.Context) error {
	request := new(UserRegistrationRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("7stheq86csp495q9u0i71rm4la"),
		Password: aws.String(request.Password),
		Username: aws.String(request.Username),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String(request.GivenName),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(request.FamilyName),
			},
			{
				Name:  aws.String("gender"),
				Value: aws.String(request.Gender),
			},
			{
				Name:  aws.String("birthdate"),
				Value: aws.String(request.Birthdate),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(request.PhoneNumber),
			},
			{
				Name:  aws.String("address"),
				Value: aws.String(request.Address),
			},
		},
	}

	_, err := config.CognitoClient.SignUp(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": awsErr.Message(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "An error occurred while registering the user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User successfully registered",
	})
}

func LoginUser(c echo.Context) error {
	// Implement user login logic with AWS Cognito
	return c.String(http.StatusOK, "User logged in")
}

func LogoutUser(c echo.Context) error {
	// Implement user logout logic with AWS Cognito
	return c.String(http.StatusOK, "User logged out")
}