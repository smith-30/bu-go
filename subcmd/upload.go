package subcmd

import (
	"sync"

	"github.com/smith-30/bu-go/helper/env"
	"github.com/smith-30/bu-go/services/logger"
	"github.com/smith-30/bu-go/services/time_util"
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
		cmd := &mysql.Command{
			DumpDst:     env.GetEnvStr("DUMP_DST", ""),
			Container:   env.GetEnvStr("DB_CONTAINER", ""),
			DumpCmd:     env.GetEnvStr("DUMP_CMD", ""),
			UserCmd:     env.GetEnvStr("DB_USER", ""),
			PasswordCmd: env.GetEnvStr("DB_PASS", ""),
			DBName:      env.GetEnvStr("DB_NAME", ""),
		}

		uc := mysql.New()
		if err := uc.Exec(cmd); err != nil {
			sl.Errorf("%v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := &folder.Command{
			Src: env.GetEnvStr("COPY_SRC", ""),
			Dst: env.GetEnvStr("COPY_DST", ""),
		}

		uc := folder.New()
		if err := uc.Exec(cmd); err != nil {
			sl.Errorf("%v", err)
		}
	}()

	wg.Wait()

	zipName := time_util.AddTS(env.GetEnvStr("ZIP_NAME", ""))
	zipPath := env.GetEnvStr("ZIP_DST", "") + zipName

	cmd := &upload.Command{
		Src:     env.GetEnvStr("BACKUP_DIR", ""),
		ZipName: zipPath,
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
