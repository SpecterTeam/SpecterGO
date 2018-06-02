package command

import (
	"errors"
	"bufio"
	"os"
	"strings"
	"github.com/SpecterTeam/SpecterGO/utils"
)

type CommandReader struct {
	lastCommand string
}

var (
	unknownCommand = errors.New("unknown command, please try /help to get a list of valid commands")
)

func NewCommandReader() CommandReader {
	c := CommandReader{}

	return c
}

func (c *CommandReader) ReadConsole() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t := strings.Replace(scanner.Text(), "/", "", 1)
			if len(t) == 0 {
				utils.HandleError(unknownCommand)
			} else {
				//execute command
			}
		}
	}()
}