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

// GetIO No input can be extracted when ran using pipe
func GetIO() (string, string, error) {
	var input string
	var output string

	if len(os.Args) > 1 && os.Args[1] == "--" {
		cmdArgs := os.Args[2:]
		input = strings.Join(cmdArgs, " ")

		if len(cmdArgs) == 0 {
			return "", "", errors.New("please provide a command")
		}

		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		outputBytes, err := cmd.CombinedOutput()

		if err != nil {
			return input, "", errors.New(fmt.Sprintf("an error occured while executing the provided command : %s", err))
		}

		output = string(outputBytes)
	} else {
		stdin, err := readCommandFromStdin()
		if err != nil {
			return "", "", err
		}

		output = strings.Join(stdin, "\n")
	}

	log.Print(output)

	return input, output, nil
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
