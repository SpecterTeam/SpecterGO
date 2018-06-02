package _default

import (
	"github.com/SpecterTeam/SpecterGO/command"
	"github.com/SpecterTeam/SpecterGO/utils"
)

func HelpCommand() command.Command {
	return command.NewCommand("help", []string{"?"}, "sends a list of commands", "specter.help", execute)
}

func execute(args []string){
	logger := utils.Logger{}
	logger.Info("Commands :")
	for name,c := range command.Commands() {
		logger.Info("/" + name + ": " + c.Description())
	}
}