package command

type Command struct {
	name        string
	aliases     []string
	description string
	permission  string
	function func([]string)
}

func NewCommand(name string, aliases []string, description string, permission string, function func([]string)) Command {
	c := Command{}
	c.SetName(name)
	c.SetAliases(aliases)
	c.SetDescription(description)
	c.SetPermission(permission)
	c.SetFunction(function)
	return c
}

func (c *Command) Function() func([]string) {
	return c.function
}

func (c *Command) SetFunction(function func([]string)) {
	c.function = function
}

func (c *Command) Permission() string {
	return c.permission
}

func (c *Command) SetPermission(permission string) {
	c.permission = permission
}

func (c *Command) Description() string {
	return c.description
}

func (c *Command) SetDescription(description string) {
	c.description = description
}

func (c *Command) Aliases() []string {
	return c.aliases
}

func (c *Command) SetAliases(aliases []string) {
	c.aliases = aliases
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) SetName(name string) {
	c.name = name
}
