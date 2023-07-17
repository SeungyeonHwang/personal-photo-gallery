package services

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	CognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (us *UserServiceImpl) RegisterUser(request *models.UserRegistrationRequest, clientId string) error {
	logrus.Infof("User %s attempting to register", request.Username)

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
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

	_, err := us.CognitoClient.SignUp(input)
	if err != nil {
		return err
	}

	logrus.Infof("User %s registered successfully", request.Username)
	return nil
}

func (us *UserServiceImpl) ConfirmUser(code string, username string, clientId string) error {
	logrus.Infof("User %s attempting to register", username)

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(clientId),
		ConfirmationCode: aws.String(code),
		Username:         aws.String(username),
	}

	_, err := us.CognitoClient.ConfirmSignUp(input)
	return err
}

func (us *UserServiceImpl) LoginUser(request *models.UserLoginRequest, clientId string) (*string, error) {
	logrus.Infof("User %s attempting to log in", request.Username)

	authFlow := "USER_PASSWORD_AUTH"
	authParam := map[string]*string{
		"USERNAME": aws.String(request.Username),
		"PASSWORD": aws.String(request.Password),
	}

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       &authFlow,
		AuthParameters: authParam,
		ClientId:       aws.String(clientId),
	}

	result, err := us.CognitoClient.InitiateAuth(input)
	if err != nil {
		return nil, err
	}

	logrus.Infof("User %s logged in successfully", request.Username)
	return result.AuthenticationResult.AccessToken, nil
}

func (us *UserServiceImpl) LogoutUser(request *models.UserGlobalLogoutRequest) error {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(request.AccessToken),
	}

	_, err := us.CognitoClient.GlobalSignOut(input)
	return err
}
