package manager

import (
	"errors"
	"github.com/MaximBayurov/task-tracker-cli/internal/command"
	"github.com/MaximBayurov/task-tracker-cli/internal/command/concrete"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"strings"
)

var instance *Manager

// Init create the instance of Manager with all allowed commands. Note: help command add after all
func Init() {
	instance = new(Manager)

	instance.commands = make(map[string]command.Executable)

	instance.commands["add"] = new(concrete.AddCommand)
	instance.commands["update"] = new(concrete.UpdateCommand)
	instance.commands["delete"] = new(concrete.DeleteCommand)

	for k, am := range mark.AllowedMarks {
		markTask := new(concrete.MarkCommand)
		markTask.NewMark = am
		key := strings.Join([]string{"mark", k}, "-")
		instance.commands[key] = markTask
	}

	instance.commands["list"] = new(concrete.ListCommand)
	helpCmd := new(concrete.HelpCommand)
	helpCmd.SetCommands(instance.commands)
	instance.commands["help"] = helpCmd
}

// GetInstance get instance of Manager
func GetInstance() (*Manager, error) {
	if instance == nil {
		return instance, errors.New("manager instance don't initialized")
	}
	return instance, nil
}
