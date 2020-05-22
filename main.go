// map is a Go reimplementation of https://github.com/soveran/map
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("usage: %s <variable> <command>", os.Args[0])
	}
	envName, command := os.Args[1], os.Args[2]
	if strings.ContainsAny(envName, "= ") {
		return fmt.Errorf("variable name %q cannot contain space of =", envName)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(bytes.TrimSpace(scanner.Bytes())) == 0 {
			continue
		}
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Env = append(os.Environ(), envName+"="+scanner.Text())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return scanner.Err()
}
