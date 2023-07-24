package models

import "mime/multipart"

type PhotoUploadRequest struct {
	FileName string
	File     *multipart.Part
}
