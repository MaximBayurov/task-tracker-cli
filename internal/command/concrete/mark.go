package concrete

import (
	"errors"
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"strconv"
)

// MarkCommand implements the processing of the mark-* command (mark-done, etc.)
type MarkCommand struct {
	command.Command
	NewMark mark.Mark
}

func (c MarkCommand) Validator(params args.Params) validator.Validator {
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

func (c MarkCommand) Handle(params args.Params) error {
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
	row.Status = c.NewMark

	err = s.Update(ID, row)
	if err != nil {
		return err
	}
	fmt.Printf("Command's mark updated successfully (ID: %v)\n", ID)
	return s.Save()
}

func (c MarkCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "Mark",
		Short: "Marking a command",
		Long:  "Pass new status with command separated with \"-\". For example, mark-in-progress change status to in-progress.",
	}
	return info, nil
}
