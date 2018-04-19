package uploader

import (
	"errors"

	"github.com/smith-30/bu-go/services/uploader/google_drive"
)

const (
	googleDrive = "googleDrive"
)

var (
	ErrInvalidName = errors.New("Invalid uploader name")
)

type Uploader interface {
	Upload(src, dst string) error
}

func NewUploader(c UploaderConfig) (Uploader, error) {
	switch c.Name {
	case googleDrive:
		return google_drive.NewUploader(c), nil
	}

	return nil, ErrInvalidName
}
