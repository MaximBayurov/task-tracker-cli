package handlers_test

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"testing"
)

type IntegerTestData struct {
	Params    args.Params
	Index     int64
	WithError bool
}

func TestIntegerRun(t *testing.T) {
	var h handlers.IntegerHandler
	for i, data := range getIntegerTestData() {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			h = handlers.IntegerHandler{}
			h.SetIndex(data.Index)
			err := h.Run(data.Params)
			if err != nil && !data.WithError {
				t.Error("Unexpected validation error")
			}
			if err == nil && data.WithError {
				t.Error("Miss expected validation error")
			}
		})
	}
}

func getIntegerTestData() []IntegerTestData {
	return []IntegerTestData{
		{args.Params{"0"}, 0, false},
		{args.Params{"0", "1", "test", "3"}, 2, true},
		{args.Params{"0"}, 2, true},
	}
}
