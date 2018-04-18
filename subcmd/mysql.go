package subcmd

type Mysql struct{}

func (f *Mysql) Help() string {
	return "backup mysql"
}

func (f *Mysql) Run(args []string) int {
	// sl := logger.GetSugaredLogger()

	return 0
}

func (f *Mysql) Synopsis() string {
	return "backup mysql"
}
