package google_drive

import "github.com/smith-30/bu-go/services/uploader"

type Uploader struct {
}

func NewUploader(c uploader.UploaderConfig) Uploader {
	return &Uploader{}
}

func (u *Uploader) Upload(src, dst string) error {

}
