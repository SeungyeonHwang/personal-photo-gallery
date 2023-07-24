package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sirupsen/logrus"
)

type CognitoService struct {
	Client    *cognitoidentityprovider.CognitoIdentityProvider
	SSMClient *ssm.SSM
	Log       *logrus.Logger
}

type S3Service struct {
	Uploader  *s3manager.Uploader
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

func NewS3Service() *S3Service {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	// Create session
	sess, err := session.NewSession()
	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatal("Failed to create session")
	}

	// Create S3 uploader
	uploader := s3manager.NewUploader(sess)
	ssmClient := ssm.New(sess)

	return &S3Service{
		Uploader:  uploader,
		SSMClient: ssmClient,
		Log:       log,
	}
}
