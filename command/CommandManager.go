package command

import (
	"github.com/SpecterTeam/SpecterGO/utils"
)

var (
	commandMap CommandMap
)

type CommandMap struct {
	commands map[string]Command
	aliases map[string]string
}

func (c *CommandMap) Commands() map[string]Command {
	return c.commands
}

func (c *CommandMap) SetCommands(commands map[string]Command) {
	c.commands = commands
}


func (c *CommandMap) AddCommand(cmd Command) {
	if c.commands == nil {
		c.commands = make(map[string]Command)
	}
	c.commands[cmd.Name()] = cmd
	c.AddAliases(cmd.Aliases(), cmd.Name())
}

func (c *CommandMap) AddAliases(aliases []string, name string)  {
	if c.aliases == nil {
		c.aliases = make(map[string]string)
	}
	for _,alias := range aliases {
		c.aliases[alias] = name
	}
}

func (c *CommandMap) CommandExist(name string) bool {
	_,exist := c.commands[name]
	_,exist2 := c.aliases[name]
	if exist == true || exist2 == true {
		return true
	} else {
		return false
	}
}

func (c *CommandMap) Command(name string) *Command {
	_,exist := c.commands[name]
	cmd := Command{}
	if exist == true {
		cmd = c.commands[name]
	} else {
		cmd = c.commands[c.aliases[name]]
	}
	return &cmd
}

func (c *CommandMap) InitCommands() {
	c.AddCommand(HelpCommand()) //register help command.
}

func GetCommandMap() *CommandMap {
	return &commandMap
}

func SendCommand(cmd string, args map[int]string) {
	if GetCommandMap().CommandExist(cmd) {
		GetCommandMap().Command(cmd).ExecuteFunction(utils.MapToArray(args))
	} else {
		utils.HandleError(UnknownCommand)
	}
}