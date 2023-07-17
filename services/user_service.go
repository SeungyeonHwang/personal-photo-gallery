package services

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/models"
)

type UserService interface {
	RegisterUser(request *models.UserRegistrationRequest, clientId string) error
	ConfirmUser(code string, username string, clientId string) error
	LoginUser(request *models.UserLoginRequest, clientId string) (*string, error)
	LogoutUser(request *models.UserGlobalLogoutRequest) error
}
