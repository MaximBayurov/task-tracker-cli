package manager

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
)

// Manager contains all allowed commands in program
type Manager struct {
	commands map[string]command.Executable
}

// HandleFromArgs - abstract command handling
func (m *Manager) HandleFromArgs(a args.Arguments) error {
	c, ok := m.commands[a.Cmd]
	if !ok {
		c, ok = m.commands["help"]
		if !ok {
			return fmt.Errorf("command \"%v\" not found", a.Cmd)
		}
	}

	v := c.Validator(a.Params)

	errs := v.Run()
	if len(errs) > 0 {
		fmt.Println("Command not executed in case of validation errors:")
		for _, err := range errs {
			fmt.Println(err)
		}
		return nil
	} else {
		return c.Handle(a.Params)
	}
}
