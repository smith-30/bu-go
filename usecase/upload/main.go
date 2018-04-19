package upload

import (
	"github.com/smith-30/bu-go/services/compress/zipper"
)

type Usecase interface {
	Exec(cmd *Command) error
}

type uc struct{}

func New() Usecase {
	return &uc{}
}

func (u *uc) Exec(cmd *Command) error {
	if err := zipper.CreateZip(cmd.Dst, cmd.ZipDst); err != nil {
		return err
	}

	// upload

	return nil
}
