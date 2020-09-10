package interpreter

import "fmt"

type Interpreter struct {
	cmdMap map[string]ICmd
}

func (object *Interpreter) RunCmd(row, column int, cmd string, args []string) (err error) {
	err = object.cmdMap[cmd].Run(args)
	if nil != err {
		err = &Error{Row: row, Column: column, E: err}
		return
	}
	fmt.Println()
	return
}

func (object *Interpreter) Interpret(raw []byte) (err error) {
	var cmd string
	var args []string
	var row int
	var column int
	var char []byte
	for 0 < len(raw) {
		ch := raw[column]
		switch ch {
		case ' ', '\t', '\r':
			if _, ok := object.cmdMap[string(char)]; ok {
				cmd = string(char)
				if 0 < len(char) {
					char = make([]byte, 0)
				}
			} else if 0 < len(char) {
				args = append(args, string(char))
				char = make([]byte, 0)
			}
			column++
			continue
		case '\n':
			row++
			if 0 < len(char) {
				if 0 >= len(cmd) {
					cmd = string(char)
				} else {
					args = append(args, string(char))
				}
				char = make([]byte, 0)
			}
			if 0 < len(cmd) {
				if err = object.RunCmd(row, column, cmd, args); nil != err {
					return
				}
				cmd = ""
				if 0 < len(args) {
					args = make([]string, 0)
				}
				if 0 < len(char) {
					char = make([]byte, 0)
				}
			}
			raw = raw[func() int {
				if column+1 < len(raw) {
					return column + 1
				}
				return len(raw)
			}():]
			column = 0
		default:
			char = append(char, ch)
			column++
			//fmt.Println(string(char))
		}
	}
	return
}

func NewInterpreter(cmdList ...ICmd) *Interpreter {
	object := &Interpreter{cmdMap: make(map[string]ICmd, 0)}
	for _, cmd := range cmdList {
		object.cmdMap[cmd.GetCmd()] = cmd
	}
	return object
}
