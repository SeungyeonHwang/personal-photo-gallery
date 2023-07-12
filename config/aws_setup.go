package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var CognitoClient *cognitoidentityprovider.CognitoIdentityProvider

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	if err != nil {
		panic(err)
	}

	CognitoClient = cognitoidentityprovider.New(sess)
}
