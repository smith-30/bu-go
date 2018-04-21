package upload

import (
	"github.com/smith-30/bu-go/services/compress/zipper"
	"github.com/smith-30/bu-go/services/uploader"
)

type Usecase interface {
	Exec(cmd *Command) error
}

type uc struct{}

func New() Usecase {
	return &uc{}
}

func (uc *uc) Exec(cmd *Command) error {
	if err := zipper.CreateZip(cmd.Src, cmd.ZipPath); err != nil {
		return err
	}

	// upload
	u, err := uploader.NewUploader(uploader.UploaderConfig{
		Name:  cmd.UseUploader,
		Token: cmd.DropboxToken,
	})
	if err != nil {
		return err
	}

	if err := u.Upload(cmd.ZipPath, cmd.ZipName); err != nil {
		return err
	}

	return nil
}
