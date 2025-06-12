package handlers_test

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"testing"
)

type MarkTestData struct {
	Params    args.Params
	Index     int64
	WithError bool
}

func TestMarkRun(t *testing.T) {
	var h handlers.MarkHandler
	for i, data := range getMarkTestData() {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			h = handlers.MarkHandler{}
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

func getMarkTestData() []MarkTestData {
	return []MarkTestData{
		{args.Params{"todo"}, 0, false},        //1
		{args.Params{"in-progress"}, 0, false}, //2
		{args.Params{"canceled"}, 0, false},    //3
		{args.Params{"done"}, 0, false},        //4
		{args.Params{"40"}, 0, true},           //5
		{args.Params{"10"}, 0, true},           //6
		{args.Params{}, 0, true},               //7
	}
}
