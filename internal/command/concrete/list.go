package concrete

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"time"
)

// ListCommand implements the processing of the task list command
type ListCommand struct {
	command.Command
}

func (c ListCommand) Validator(params args.Params) validator.Validator {
	v := validator.Validator{}
	v.SetParams(params)

	cnt := handlers.CountHandler{}
	cnt.SetTo(1)
	v.AddHandler(&cnt)

	if len(params) == 1 {
		mrk := handlers.MarkHandler{}
		mrk.SetIndex(0)
		v.AddHandler(&mrk)
	}

	return v
}

func (c ListCommand) Handle(params args.Params) error {
	s, err := storage.GetInstance()
	if err != nil {
		return err
	}

	var rows []storage.DataRow
	if len(params) > 0 {
		status, err := mark.FromString(params[0])
		if err != nil {
			return err
		}
		rows = s.GetByStatus(status)
	} else {
		rows = s.GetAll()
	}

	rowTmpl := "%5v | %50v | %15v | %25v | %25v \n"
	fmt.Printf(
		rowTmpl,
		"ID", "Description", "Status", "CreatedAt", "UpdatedAt",
	)
	for _, row := range rows {
		fmt.Printf(
			rowTmpl,
			row.ID,
			row.Description,
			row.Status.ToString(),
			row.CreatedAt.Format(time.DateTime),
			row.UpdatedAt.Format(time.DateTime),
		)
	}

	return nil
}

func (c ListCommand) Info() (command.Info, error) {
	info := command.Info{
		Name:  "List",
		Short: "Listing all tasks",
		Long:  "Listing all tasks with status filter",
	}
	return info, nil
}
