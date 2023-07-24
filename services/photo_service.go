// [photo_service.go]
package services

import (
	"github.com/SeungyeonHwang/personal-photo-gallery/models"
)

type PhotoService interface {
	UploadPhoto(request *models.PhotoUploadRequest, bucketName string) error
}
