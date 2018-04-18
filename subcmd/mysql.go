package subcmd

import (
	"go.uber.org/zap"
)

type Mysql struct{}

func (f *Boot) Help() string {
	return "backup mysql"
}

func (f *Boot) Run(args []string) int {
	logger, _ := zap.NewDevelopment()

	return 0
}

func (f *Boot) Synopsis() string {
	return "backup mysql"
}
