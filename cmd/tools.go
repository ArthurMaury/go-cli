package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func readClean() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(s, "\r\n"), nil
}

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

func contains(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
