package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sirupsen/logrus"
)

type CognitoService struct {
	Client    *cognitoidentityprovider.CognitoIdentityProvider
	SSMClient *ssm.SSM
	Log       *logrus.Logger
}

func NewCognitoService() *CognitoService {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	// Create session
	sess, err := session.NewSession()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatal("Failed to create session")
	}

	// Create clients
	cognitoClient := cognitoidentityprovider.New(sess)
	ssmClient := ssm.New(sess)

	return &CognitoService{
		Client:    cognitoClient,
		SSMClient: ssmClient,
		Log:       log,
	}
}
