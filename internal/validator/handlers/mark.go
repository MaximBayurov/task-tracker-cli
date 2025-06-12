package handlers

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
)

type MarkHandler struct {
	index int64
}

func (h *MarkHandler) SetIndex(index int64) {
	h.index = index
}

func (h *MarkHandler) Run(params args.Params) error {
	if h.index+1 > int64(len(params)) {
		return fmt.Errorf("param with index %v isn't exist", h.index)
	}
	code := params[h.index]
	_, err := mark.FromString(code)
	if err != nil {
		return fmt.Errorf("the mark with code \"%v\" doesn't exist", code)
	}
	return nil
}
