package concrete

import (
	"errors"
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"strconv"
)

// UpdateCommand implements the processing of the update command
type UpdateCommand struct {
	command.Command
}

func (c UpdateCommand) Validator(params args.Params) validator.Validator {
	v := validator.Validator{}
	v.SetParams(params)

	cnt := handlers.CountHandler{}
	cnt.SetFrom(2)
	cnt.SetTo(2)
	v.AddHandler(&cnt)

	integer := handlers.IntegerHandler{}
	integer.SetIndex(0)
	v.AddHandler(&integer)

	tsk := handlers.TaskHandler{}
	tsk.SetIndex(0)
	v.AddHandler(&tsk)

	return v
}

func (c UpdateCommand) Handle(params args.Params) error {
	if len(params) > 2 {
		return errors.New("not enough params, at least two require")
	}
	s, err := storage.GetInstance()
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	ID := int64(i)
	row, err := s.GetById(ID)
	if err != nil {
		return err
	}
	row.Description = params[1]

	err = s.Update(ID, row)
	if err != nil {
		return err
	}
	fmt.Printf("Command updated successfully (ID: %v)\n", ID)
	return s.Save()
}

func (c UpdateCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "Update",
		Short: "Updating a command by ID",
		Long:  "",
	}
	return info, nil
}
