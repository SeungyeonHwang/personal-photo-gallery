package services

type SSMService interface {
	GetParameter(name string) (string, error)
}
