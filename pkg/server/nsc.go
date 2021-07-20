package server

import (
	"bytes"
	"fmt"
	"os/exec"
)

func NscAddUser(name string, options ...string) error {
	// TODO
	fmt.Print(name, options)

	return nil
}

func NscAddOperator(name string, options ...string) error {
	// TODO
	fmt.Print(name, options)

	return nil
}

func runNscCommand(args ...string) error {
	cmd := exec.Command("nsc", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err == nil {
		fmt.Printf("Output: %s", out.String())
	}
	return err
}
