package SpecterGO

import (
	"github.com/SpecterTeam/SpecterGO/command"
	"github.com/SpecterTeam/SpecterGO/utils"
	"strconv"
)

func InitCommands() {
	command.GetCommandMap().AddCommand(HelpCommand())
	command.GetCommandMap().AddCommand(StopCommand())
}

func HelpCommand() command.Command {
	return command.NewCommand("help", []string{"?"}, "sends a list of commands", "specter.help", func(args []string){
		GetLogger().Info("Commands :")
		if len(args) == 0 {
			for name, c := range command.GetCommandMap().Commands() {
				GetLogger().Info("/" + name + ": " + c.Description())
			}
		} else {
			page,err := strconv.Atoi(args[0])
			if err != nil {
				utils.HandleError(err)
			} else {
				list := makeHelpList()
				max := len(list)
				if page > max {
					for name,c := range list[0] {
						GetLogger().Info("/" + name + ": " + c.Description())
					}
				} else {
					page--
					for name,c := range list[page] {
						GetLogger().Info("/" + name + ": " + c.Description())
					}
				}
			}
		}
	})
}

func StopCommand() command.Command {
	return command.NewCommand("stop", []string{"shutdown", "exit"}, "Stop the server", "specter.stop", func(args []string) {
	GetLogger().Info("Stopping the server...")
	GetServer().Shutdown()
	})
}


func makeHelpList() map[int]map[string]command.Command {
	list := make(map[int]map[string]command.Command)
	count,page := 0,0
	for name,c := range command.GetCommandMap().Commands() {
		if count == 5 {
			page++
		}
		list[page][name] = c
		count++
	}

	return list
}
