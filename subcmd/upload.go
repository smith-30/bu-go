package subcmd

import (
	"sync"

	"github.com/smith-30/bu-go/services/logger"
	"github.com/smith-30/bu-go/usecase/folder"
	"github.com/smith-30/bu-go/usecase/mysql"
	"github.com/smith-30/bu-go/usecase/upload"
)

type Upload struct{}

func (u *Upload) Help() string {
	return "upload command"
}

func (u *Upload) Run(args []string) int {
	sl := logger.GetSugaredLogger()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := &mysql.Command{}

		uc := mysql.New()
		if err := uc.Exec(cmd); err != nil {
			sl.Errorf("%v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := &folder.Command{
			Src: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/src",
			Dst: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied",
			// ZipName: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied.zip",
		}

		uc := folder.New()
		if err := uc.Exec(cmd); err != nil {
			sl.Errorf("%v", err)
		}
	}()

	wg.Wait()

	cmd := &upload.Command{
		Src:     "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/src",
		ZipName: "/Users/kouhei/go/src/github.com/smith-30/bu-go/tmp/copied.zip",
	}

	uc := upload.New()
	if err := uc.Exec(cmd); err != nil {
		sl.Errorf("%v", err)
		return 1
	}

	return 0
}

func (u *Upload) Synopsis() string {
	return "upload"
}
