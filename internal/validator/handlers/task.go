package handlers

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/storage"
	"strconv"
)

type TaskHandler struct {
	index int64
}

func (h *TaskHandler) SetIndex(index int64) {
	h.index = index
}

func (h *TaskHandler) Run(params args.Params) error {
	if h.index+1 > int64(len(params)) {
		return fmt.Errorf("param with index %v isn't exist", h.index)
	}

	IDraw := params[h.index]
	validationErr := fmt.Errorf("the task with ID=\"%v\" doesn't exist", IDraw)

	ID, err := strconv.Atoi(IDraw)
	if err != nil {
		return validationErr
	}

	s, err := storage.GetInstance()
	if err != nil {
		return validationErr
	}
	_, err = s.GetById(int64(ID))
	if err != nil {
		return validationErr
	}
	return nil
}
