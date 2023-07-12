package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

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
