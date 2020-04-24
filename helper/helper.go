package helper

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Pf is Sprintf with runtime function
func Pf(message string, a ...interface{}) string {
	pc, _, _, _ := runtime.Caller(1)
	s := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	newMessage := fmt.Sprintf("[%s] %s", s[len(s)-1], message)
	return fmt.Sprintf(newMessage, a...)
}

// Pef is Sprintf with runtime function and return error
func Pef(message string, a ...interface{}) error {
	pc, _, _, _ := runtime.Caller(1)
	s := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	newMessage := fmt.Sprintf("[%s] %s", s[len(s)-1], message)
	return fmt.Errorf(newMessage, a...)
}

// Pe is wrap run function with error
func Pe(err error) error {
	pc, _, _, _ := runtime.Caller(1)
	s := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	newMessage := fmt.Sprintf("[%s] %s", s[len(s)-1], err.Error())
	return errors.New(newMessage)
}
