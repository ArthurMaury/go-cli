package cmd

import (
	"bufio"
	"os"
	"strings"
)

func readClean() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(s, "\r\n"), nil
}
