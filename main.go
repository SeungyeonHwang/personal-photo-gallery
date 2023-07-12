package main

import (
	"net/http"

	"github.com/SeungyeonHwang/personal-photo-gallery/handlers"
	"github.com/SeungyeonHwang/personal-photo-gallery/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.JWTValidationMiddleware)

	// home
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// userHandler
	e.POST("/users/register", handlers.RegisterUser)
	e.POST("/users/login", handlers.LoginUser)
	e.POST("/users/confirm", handlers.ConfirmUser)
	e.GET("/users/logout", handlers.LogoutUser)

	// photoHandler
	e.POST("/photos", handlers.UploadPhoto)
	e.GET("/photos", handlers.GetPhotoList)
	e.GET("/photos/:id", handlers.GetPhotoDetails)
	e.PUT("/photos/:id", handlers.UpdatePhotoDetails)
	e.DELETE("/photos/:id", handlers.DeletePhoto)

	e.Start(":8080")
}
