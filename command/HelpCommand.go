package command

import (
	"github.com/SpecterTeam/SpecterGO/utils"
)

func HelpCommand() Command {
	return NewCommand("help", []string{"?"}, "sends a list of commands", "specter.help", executeHelp)
}

func executeHelp(args []string){
	logger := utils.Logger{}
	logger.Info("Commands :")
	for name,c := range GetCommandMap().Commands() {
		logger.Info("/" + name + ": " + c.Description())
	}
}