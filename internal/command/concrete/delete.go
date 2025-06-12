package concrete

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"strconv"
)

// DeleteCommand implements the processing of the task delete command
type DeleteCommand struct {
	command.Command
}

func (c DeleteCommand) Validator(params args.Params) validator.Validator {
	v := validator.Validator{}
	v.SetParams(params)

	cnt := handlers.CountHandler{}
	cnt.SetFrom(1)
	cnt.SetTo(1)
	v.AddHandler(&cnt)

	integer := handlers.IntegerHandler{}
	integer.SetIndex(0)
	v.AddHandler(&integer)

	tsk := handlers.TaskHandler{}
	tsk.SetIndex(0)
	v.AddHandler(&tsk)

	return v
}

func (c DeleteCommand) Handle(params args.Params) error {
	s, err := storage.GetInstance()
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	ID := int64(i)
	err = s.Delete(ID)
	if err != nil {
		return err
	}

	fmt.Printf("Command deleted successfully (ID: %v)\n", ID)
	return s.Save()
}

func (c DeleteCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "Delete",
		Short: "Delete a command by ID",
		Long:  "",
	}
	return info, nil
}
