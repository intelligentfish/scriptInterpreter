package main

import (
	"flag"
	"io/ioutil"

	"interpreter/interpreter"
)

func main() {
	var scriptPath string
	flag.StringVar(&scriptPath, "script_path", "", "specify script path")
	flag.Parse()
	if 0 >= len(scriptPath) {
		flag.Usage()
		return
	}
	input, err := ioutil.ReadFile(scriptPath)
	if nil != err {
		panic(err)
	}
	interpreter.NewInterpreter(
		&interpreter.Now{},
		&interpreter.Sleep{},
		&interpreter.Echo{},
		&interpreter.Get{},
	).Interpret(input)
}
