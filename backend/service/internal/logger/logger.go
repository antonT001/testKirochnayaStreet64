package logger

import (
	"fmt"
)

type Logger interface {
	Log(v ...interface{})
	Panic(v ...interface{})
}

type logger struct {
}

func New() Logger {
	return &logger{}
}

func (l *logger) Log(v ...interface{}) {
	fmt.Println(v...)
}

func (l *logger) Panic(v ...interface{}) {
	fmt.Println(v...)
	panic(struct{}{})
}
