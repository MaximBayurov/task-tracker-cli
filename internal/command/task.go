package command

import (
	"errors"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
)

// Executable implement it to extend program commands' availability
type Executable interface {
	Handle(params args.Params) error
	Info() (Info, error)
	Validator(params args.Params) validator.Validator
}

// Command base struct for commands
type Command struct {
}

// Handle method for command handling
func (c Command) Handle(params args.Params) error {
	return errors.New("handle method not implemented")
}

// Info returns data about command
func (c Command) Info() (Info, error) {
	return Info{}, errors.New("info method not implemented")
}

// Validator returns the command parameter validator
func (c Command) Validator(params args.Params) validator.Validator {
	return validator.Validator{}
}
