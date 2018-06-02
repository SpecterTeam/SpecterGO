package command

import "github.com/SpecterTeam/SpecterGO/command/default"

var (
	commands map[string]Command
)

func Commands() map[string]Command {
	return commands
}

func AddCommand(c Command) {
	Commands()[c.Name()] = c
}

func InitCommands() {
	_default.HelpCommand() //register help command.
}