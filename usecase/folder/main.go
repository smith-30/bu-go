package folder

import "github.com/otiai10/copy"

type Usecase interface {
	Exec(cmd *Command) error
}

type uc struct{}

var Copier = func(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}

func New() Usecase {
	return &uc{}
}

func (u *uc) Exec(cmd *Command) error {
	if err := Copier(cmd.Src, cmd.Dst); err != nil {
		return err
	}

	return nil
}
