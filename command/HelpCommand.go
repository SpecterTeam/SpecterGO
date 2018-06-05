package command

import (
	"github.com/SpecterTeam/SpecterGO/utils"
	"strconv"
)

func HelpCommand() Command {
	return NewCommand("help", []string{"?"}, "sends a list of commands", "specter.help", executeHelp)
}

func executeHelp(args []string){
	logger := utils.Logger{}
	logger.Info("Commands :")
	if len(args) == 0 {
		for name, c := range GetCommandMap().Commands() {
			logger.Info("/" + name + ": " + c.Description())
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
					logger.Info("/" + name + ": " + c.Description())
				}
			} else {
				page--
				for name,c := range list[page] {
					logger.Info("/" + name + ": " + c.Description())
				}
			}
		}
	}
}

func makeHelpList() map[int]map[string]Command {
	list := make(map[int]map[string]Command)
	count,page := 0,0
	for name,c := range GetCommandMap().Commands() {
		if count == 5 {
			page++
		}
		list[page][name] = c
		count++
	}

	return list
}