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
		"dockernginxproxyphp_app_mysql_1",
		"/usr/bin/mysqldump",
		"-uroot",
		"--password=XPQ_N0+0dIm",
		"app_mysql",
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

	err = ioutil.WriteFile("./out.sql", bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
