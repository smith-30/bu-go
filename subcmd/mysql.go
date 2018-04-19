package subcmd

import (
	"github.com/smith-30/bu-go/services/logger"
	"github.com/smith-30/bu-go/usecase/mysql"
)

type Mysql struct{}

func (f *Mysql) Help() string {
	return "backup mysql"
}

func (f *Mysql) Run(args []string) int {
	sl := logger.GetSugaredLogger()

	cmd := &mysql.Command{}

	uc := mysql.New()
	if err := uc.Exec(cmd); err != nil {
		sl.Fatalf("%v\n", err)
		return 1
	}
	return 0
}

func (f *Mysql) Synopsis() string {
	return "backup mysql"
}
