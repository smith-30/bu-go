package mysql

type Command struct {
	DumpDst,
	Container,
	DumpCmd,
	UserCmd,
	PasswordCmd,
	DBName string
}
