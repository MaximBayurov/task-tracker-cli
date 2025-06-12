package concrete

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
)

// AddCommand implements the processing of the task add command
type AddCommand struct {
	command.Command
}

func (c AddCommand) Validator(params args.Params) validator.Validator {
	v := validator.Validator{}
	v.SetParams(params)

	cnt := handlers.CountHandler{}
	cnt.SetFrom(1)
	cnt.SetTo(1)
	v.AddHandler(&cnt)

	return v
}

func (c AddCommand) Handle(params args.Params) error {
	row := storage.DataRow{
		ID:          0,
		Description: params[0],
		Status:      mark.Todo,
	}
	s, err := storage.GetInstance()
	if err != nil {
		return err
	}
	fmt.Printf("Command added successfully (ID: %v)\n", s.Add(row))
	return s.Save()
}

func (c AddCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "Add",
		Short: "Adding a new command",
		Long:  "",
	}
	return info, nil
}
