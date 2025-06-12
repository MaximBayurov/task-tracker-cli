package concrete

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
)

// HelpCommand implements the processing of the help command
type HelpCommand struct {
	command.Command
	cmds map[string]command.Executable
}

func (c *HelpCommand) Validator(params args.Params) validator.Validator {
	v := validator.Validator{}
	v.SetParams(params)

	cnt := handlers.CountHandler{}
	cnt.SetFrom(0)
	cnt.SetTo(1)
	v.AddHandler(&cnt)

	return v
}

func (c *HelpCommand) Handle(params args.Params) error {
	switch len(params) {
	case 1:
		code := string(params[0])
		cmd, err := c.byCode(code)
		if err != nil {
			return err
		}
		info, err := cmd.Info()
		if err != nil {
			return err
		}
		fmt.Printf(
			"Command: %v\n"+
				"Name: %v\n"+
				"Summary: %v\n"+
				"Details: %v",
			code,
			info.Name,
			info.Short,
			info.Long,
		)
	case 0:
		fmt.Printf("%20v | %10v | %v\n", "Command", "Name", "Summary")
		for code, cmd := range c.commands() {
			info, err := cmd.Info()
			if err != nil {
				continue
			}
			fmt.Printf("%20v | %10v | %v\n", code, info.Name, info.Short)
		}
	}

	return nil
}

func (c *HelpCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "Help",
		Short: "Print help info about available commands",
		Long:  "",
	}
	return info, nil
}

func (c *HelpCommand) SetCommands(commands map[string]command.Executable) {
	c.cmds = commands
}

func (c *HelpCommand) commands() map[string]command.Executable {
	return c.cmds
}

func (c *HelpCommand) byCode(code string) (command.Executable, error) {
	cmd, ok := c.cmds[code]
	if !ok {
		return command.Command{}, fmt.Errorf("cannot find command with code \"%v\"", code)
	}
	return cmd, nil
}
