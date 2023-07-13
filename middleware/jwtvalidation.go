package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

// TODO: gonnado delete
// JWT validation could be done here or with a Custom Authorizer

// JWT 검증을 위해 API Gateway에서 Custom Authorizer를 사용합니다.

// Lambda 함수를 생성하여 JWT를 검증한 다음 이 Lambda 함수를 SAM 템플릿에서 Custom Authorizer로 설정할 수 있습니다. 이 함수는 events.APIGatewayCustomAuthorizerRequest를 매개변수로 사용하고 events.APIGatewayCustomAuthorizerResponse를 반환합니다.

// 이러한 모든 Lambda 함수는 개별적으로 패키징되어 AWS Lambda에 업로드되어야 합니다.

// 미들웨어의 경우 AWS Lambda와 API Gateway는 본질적으로 동일한 미들웨어 개념을 지원하지 않기 때문에 각 Lambda 함수에 기능을 포함하거나 다른 AWS 서비스를 사용하여 기능을 모방해야 합니다. 예를 들어 API Gateway에서 JWT 검증 미들웨어용 사용자 지정 권한 부여자로 Lambda 함수를 사용할 수 있습니다. 이 사용자 지정 권한 부여자는 기본 Lambda 함수 전에 트리거되고 JWT를 검증한 다음 검증이 성공하면 요청을 Lambda 함수로 전달합니다.

// DynamoDB 또는 다른 데이터베이스를 사용하는 경우 Lambda의 상태 비저장 특성으로 인해 연결을 다르게 관리해야 합니다.

// 또한 AWS Lambda는 요청과 응답을 처리하는 데 다른 방식을 사용하므로 Lambda 방식에 맞게 코드를 조정해야 합니다. net/http 또는 Echo의 방식을 직접 사용하여 요청 및 응답을 처리할 수 없습니다.

// 마지막으로 API Gateway 및 Lambda 함수가 제대로 통신하고 작동할 수 있도록 AWS에서 적절한 보안 설정(예: IAM 역할)을 지정하는 것을 잊지 마십시오.

func JWTValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Skip authentication for the login, registration, and confirmation endpoints
		if c.Path() == "/users/register" || c.Path() == "/users/login" || c.Path() == "/users/confirm" {
			return next(c)
		}

		// Fetch the JWKS from your Cognito User Pool
		ctx := context.Background()
		keyset, err := jwk.Fetch(ctx, "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_FdfTk7HA0/.well-known/jwks.json")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: " + err.Error()})
		}

		// Get the JWT from the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Handle error: malformed authorization header
			return echo.NewHTTPError(http.StatusUnauthorized, "malformed authorization header")
		}
		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse([]byte(tokenString), jwt.WithKeySet(keyset))
		if err != nil {
			// Handle error: invalid token
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		// Add the token to the context
		c.Set("user", token)

		// Call the next handler
		return next(c)
	}
}
