package interpreter

type ICmd interface {
	GetCmd() string
	Run(args []string) (err error)
}
