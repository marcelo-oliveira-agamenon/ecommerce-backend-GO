package ports

import "mime/multipart"

type SaveFileResponse struct {
	ImageKey string
	ImageURL string
}

type StorageService interface {
	SaveFileAWS(file multipart.File, fileHeader string, fileSize int64, typeModel string) (*SaveFileResponse, error)
	DeleteFileAWS(fileName string) (bool, error)
}
