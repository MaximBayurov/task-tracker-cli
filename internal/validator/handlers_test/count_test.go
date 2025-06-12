package handlers_test

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"testing"
)

type CountTestData struct {
	Params    args.Params
	From      int64
	To        int64
	WithError bool
}

func TestCountRun(t *testing.T) {
	var h handlers.CountHandler
	for i, data := range getCountTestData() {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			h = handlers.CountHandler{}
			if data.From > 0 {
				h.SetFrom(data.From)
			}
			if data.To > 0 {
				h.SetTo(data.To)
			}
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

func getCountTestData() []CountTestData {
	return []CountTestData{
		{args.Params{"1", "2"}, 2, 2, false},          //1
		{args.Params{"1", "2"}, 0, 2, false},          //2
		{args.Params{"1", "2"}, 2, 0, false},          //3
		{args.Params{"1"}, 2, 0, true},                //4
		{args.Params{"1", "2", "3"}, 2, 2, true},      //5
		{args.Params{"1", "2", "3"}, 3, 2, false},     //6
		{args.Params{"1", "2"}, 3, 2, false},          //7
		{args.Params{"1"}, 3, 2, true},                //8
		{args.Params{"1", "2", "3", "4"}, 3, 2, true}, //9
	}
}
