package handlers_test

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"github.com/MaximBayurov/task-tracker-cli/internal/validator/handlers"
	"strconv"
	"testing"
	"time"
)

type TaskTestData struct {
	Params    args.Params
	Index     int64
	WithError bool
}

func TestTaskRun(t *testing.T) {
	_ = storage.Init("todo-list-test.json")
	s, _ := storage.GetInstance()
	var h handlers.TaskHandler
	for i, data := range getTaskTestData(s) {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			h = handlers.TaskHandler{}
			h.SetIndex(data.Index)
			err := h.Run(data.Params)
			if err != nil && !data.WithError {
				t.Errorf("Unexpected validation error %v", data)
			}
			if err == nil && data.WithError {
				t.Error("Miss expected validation error")
			}
		})
	}
}

func getTaskTestData(s *storage.Storage) []TaskTestData {
	ID := s.Add(storage.DataRow{
		Description: "Test 1",
		Status:      mark.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	return []TaskTestData{
		{args.Params{strconv.FormatInt(ID, 10)}, 0, false},
		{args.Params{"test"}, 0, true},
		{args.Params{}, 0, true},
		{args.Params{"200"}, 0, true},
	}
}
