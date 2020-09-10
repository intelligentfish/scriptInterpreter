package interpreter

import (
	"strconv"
	"time"
)

type Sleep struct {
}

func (object *Sleep) GetCmd() string {
	return "sleep"
}

func (object *Sleep) Run(args []string) (err error) {
	if 2 != len(args) {
		err = ErrArgs
	}
	var number int64
	if number, err = strconv.ParseInt(args[0], 10, 64); nil != err {
		return
	}
	var duration time.Duration
	switch args[1] {
	case "h":
		duration = time.Duration(number) * time.Hour
	case "m":
		duration = time.Duration(number) * time.Minute
	case "s":
		duration = time.Duration(number) * time.Second
	case "ms":
		duration = time.Duration(number) * time.Millisecond
	case "ns":
		duration = time.Duration(number) * time.Nanosecond
	default:
		err = ErrArgs
	}
	if 0 < duration {
		time.Sleep(duration)
	}
	return
}
