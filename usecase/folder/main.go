package folder

import (
	"github.com/smith-30/bu-go/services/file/copier"
)

type Usecase interface {
	Exec(cmd *Command) error
}

type uc struct{}

func New() Usecase {
	return &uc{}
}

func (u *uc) Exec(cmd *Command) error {
	if err := copier.Copy(cmd.Src, cmd.Dst); err != nil {
		return err
	}

	return nil
}
