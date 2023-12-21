package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetOutput() (string, error) {
	var output string

	if len(os.Args) > 1 && os.Args[1] == "--" {
		cmdArgs := os.Args[2:]

		if len(cmdArgs) == 0 {
			return "", errors.New("please provide a command")
		}

		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		outputBytes, err := cmd.CombinedOutput()

		if err != nil {
			return "", errors.New(fmt.Sprintf("an error occured while executing the provided command : %s", err))
		}

		output = string(outputBytes)
	} else {
		stdin, err := readCommandFromStdin()
		if err != nil {
			return "", err
		}

		output = strings.Join(stdin, "\n")
	}

	log.Print(output)

	return output, nil
}

func readCommandFromStdin() ([]string, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("an error occured while reading standard input : %s", err))
	}

	return splitLines(string(data)), nil
}

func splitLines(s string) []string {
	var lines []string
	for _, line := range splitWithoutEmpty(s, '\n') {
		lines = append(lines, line)
	}
	return lines
}

func splitWithoutEmpty(s string, sep rune) []string {
	var fields []string
	for _, field := range strings.FieldsFunc(s, func(c rune) bool { return c == sep }) {
		if field != "" {
			fields = append(fields, field)
		}
	}
	return fields
}
