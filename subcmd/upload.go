package subcmd

import (
	"github.com/smith-30/bu-go/services/logger"
	"github.com/smith-30/bu-go/usecase/upload"
)

type Upload struct{}

func (u *Upload) Help() string {
	return "upload command"
}

func (u *Upload) Run(args []string) int {
	sl := logger.GetSugaredLogger()

	cmd := &upload.Command{
		Src:     "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/src",
		ZipName: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied.zip",
	}

	uc := upload.New()
	if err := uc.Exec(cmd); err != nil {
		sl.Fatalf("%v\n", err)
		return 1
	}
	return 0
}

func (u *Upload) Synopsis() string {
	return "upload"
}
