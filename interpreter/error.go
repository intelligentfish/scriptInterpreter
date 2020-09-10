package interpreter

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrKeywords = errors.New("keywords error")
	ErrArgs     = errors.New("cmd arguments error")
)

type Error struct {
	Row    int   `json:"row"`
	Column int   `json:"column"`
	E      error `json:"error"`
}

func (object *Error) Error() string {
	return fmt.Sprintf("row: %d, column: %d, error: %s",
		object.Row, object.Column, object.E)
}

func (object *Error) String() string {
	raw, _ := json.Marshal(object)
	return string(raw)
}
