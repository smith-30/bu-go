package subcmd

import (
	"github.com/smith-30/bu-go/services/logger"
	"github.com/smith-30/bu-go/usecase/folder"
)

type Folder struct{}

func (f *Folder) Help() string {
	return "backup folder"
}

func (f *Folder) Run(args []string) int {
	sl := logger.GetSugaredLogger()

	cmd := &folder.Command{
		Src: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/src",
		Dst: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied",
		// ZipName: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied.zip",
	}

	uc := folder.New()
	if err := uc.Exec(cmd); err != nil {
		sl.Fatalf("%v\n", err)
		return 1
	}

	return 0
}

func (f *Folder) Synopsis() string {
	return "backup folder"
}
