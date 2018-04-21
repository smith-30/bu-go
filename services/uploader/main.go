package uploader

import (
	"errors"

	"github.com/smith-30/bu-go/services/uploader/dropbox"
	"github.com/smith-30/bu-go/services/uploader/google_drive"
)

const (
	googleDrive = "googleDrive"
	Dropbox     = "dropbox"
)

var (
	errInvalidName = errors.New("Invalid uploader name")
)

type Uploader interface {
	Upload(src, dst string) error
}

func NewUploader(c UploaderConfig) (Uploader, error) {
	switch c.Name {
	case googleDrive:
		return google_drive.NewUploader(google_drive.Config{
			ConfPath: c.ConfPath,
		}), nil
	case Dropbox:
		return dropbox.NewUploader(dropbox.Config{
			Token: c.Token,
		}), nil
	}

	return nil, errInvalidName
}
