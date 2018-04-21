package dropbox

import (
	"os"

	d "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

type (
	Config struct {
		Token string
	}

	Uploader struct {
		Config
	}
)

func NewUploader(c Config) *Uploader {
	return &Uploader{
		c,
	}
}

func (u *Uploader) Upload(src, dst string) error {
	config := d.Config{
		Token:    u.Token,
		LogLevel: d.LogOff, // if needed, set the desired logging level. Default is off
	}

	cli := files.New(config)
	req := files.NewCommitInfo(dst)
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	_, err = cli.Upload(req, f)
	if err != nil {
		return err
	}
	return nil
}
