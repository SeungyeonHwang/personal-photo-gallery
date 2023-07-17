package models

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

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserConfirmationRequest struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}

type UserGlobalLogoutRequest struct {
	AccessToken string `json:"access_token"`
}
