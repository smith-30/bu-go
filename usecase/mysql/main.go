package mysql

import (
	"io/ioutil"
	"os/exec"
)

type Usecase interface {
	Exec(cmd *Command) error
}

type uc struct{}

func New() Usecase {
	return &uc{}
}

func (u *uc) Exec(c *Command) error {
	cmd := exec.Command(
		"docker",
		"exec",
		c.Container,
		c.DumpCmd,
		c.UserCmd,
		c.PasswordCmd,
		c.DBName,
	)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.DumpDst, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
