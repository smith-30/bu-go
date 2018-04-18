package subcmd

import (
	"go.uber.org/zap"
)

type Folder struct{}

func (f *Boot) Help() string {
	return "backup folder"
}

func (f *Boot) Run(args []string) int {
	logger, _ := zap.NewDevelopment()

	return 0
}

func (f *Boot) Synopsis() string {
	return "backup folder"
}
