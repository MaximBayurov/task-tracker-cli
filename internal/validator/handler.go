package validator

import "github.com/MaximBayurov/task-tracker-cli/internal/args"

// Handler implement it to extend validation handlers
type Handler interface {
	Run(params args.Params) error
}
