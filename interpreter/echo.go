package interpreter

import (
	"fmt"
	"strings"
)

type Echo struct {
}

func (object *Echo) GetCmd() string {
	return "echo"
}

func (object *Echo) Run(args []string) (err error) {
	_, err = fmt.Println(strings.Join(args, " "))
	return
}
