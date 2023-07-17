package main

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/config"
	"github.com/SeungyeonHwang/personal-photo-gallery/handlers"
	"github.com/SeungyeonHwang/personal-photo-gallery/services"
	"github.com/aws/aws-lambda-go/lambda"
)

// 一般的なログアウトプロセスではなく、セキュリティ的な機能としてユーザーのすべてのトークンを無効化させる機能。
// 通常のログアウトプロセスはクライアント側で行う。
func main() {
	cognitoService := config.NewCognitoService()

	userHandler := &handlers.UserHandler{
		UserService: &services.UserServiceImpl{
			CognitoClient: cognitoService.Client,
		},
	}

	lambda.Start(userHandler.HandleLogoutRequest)
}
