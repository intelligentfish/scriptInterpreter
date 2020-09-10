package interpreter

import (
	"fmt"
	"time"
)

type Now struct {
}

func (object *Now) GetCmd() string {
	return "now"
}

func (object *Now) Run(args []string) (err error) {
	fmt.Println(time.Now())
	return
}
