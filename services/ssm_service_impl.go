package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/sirupsen/logrus"
)

type SSMServiceImpl struct {
	SSMClient ssmiface.SSMAPI
}

func (ss *SSMServiceImpl) GetParameter(name string) (string, error) {
	input := &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}

	result, err := ss.SSMClient.GetParameter(input)
	if err != nil {
		logrus.Errorf("Failed to get parameter: %v", err)
		return "", err
	}

	logrus.Infof("Parameter %s retrieved successfully", name)
	return *result.Parameter.Value, nil
}
