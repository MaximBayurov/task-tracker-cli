package handlers

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"strconv"
)

type IntegerHandler struct {
	index int64
}

func (h *IntegerHandler) SetIndex(index int64) {
	h.index = index
}

func (h *IntegerHandler) Run(params args.Params) error {
	if h.index+1 > int64(len(params)) {
		return fmt.Errorf("param with index %v isn't exist", h.index)
	}
	param := params[h.index]
	_, err := strconv.Atoi(param)
	if err != nil {
		return fmt.Errorf("param with index %v isn't an integer", h.index)
	}
	return nil
}
