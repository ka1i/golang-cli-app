package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Usage() string {
	var buffer strings.Builder

	return buffer.String()
}

func Flags() (bool, error) {
	var err error
	var argv = os.Args[1:]
	var isBreak bool = false

	switch argv[0] {
	default:
		err = errors.New("please check usage")
	}
	return isBreak, err
}

func Execute() error {
	var err error

	fmt.Println("golang cli app template")

	return err
}
