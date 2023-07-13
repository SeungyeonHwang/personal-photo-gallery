package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

var CognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	if err != nil {
		panic(err)
	}

	CognitoClient = cognitoidentityprovider.New(sess)
}

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

// TODO:
// 1. 에러처리를 위한 자세한 메세지 추가
// 2. Lambda와 통합되는 AWS CloudWatch Logs를 사용하여 중요한 정보와 오류를 기록( 모니터링 및 오류 추적은 서버리스 애플리케이션의 상태와 안정성을 유지하는 데 중요)
// 3. ClineId ->  보안 환경 변수 또는 AWS Secrets Manager
// 4. models, service 에 분리 시키기? 필요다면
// 5. S3에 업로드하기 lambda
// 6. 불필요한 Dir 삭제
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	regReq := new(UserRegistrationRequest)
	err := json.Unmarshal([]byte(request.Body), regReq)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid input parameters",
		}, nil
	}

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("7stheq86csp495q9u0i71rm4la"),
		Password: aws.String(regReq.Password),
		Username: aws.String(regReq.Username),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String(regReq.GivenName),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(regReq.FamilyName),
			},
			{
				Name:  aws.String("gender"),
				Value: aws.String(regReq.Gender),
			},
			{
				Name:  aws.String("birthdate"),
				Value: aws.String(regReq.Birthdate),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(regReq.PhoneNumber),
			},
			{
				Name:  aws.String("address"),
				Value: aws.String(regReq.Address),
			},
		},
	}

	_, err = CognitoClient.SignUp(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "An error occurred while registering the user",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "User registered successfully",
	}, nil
}

func main() {
	lambda.Start(handler)
}
