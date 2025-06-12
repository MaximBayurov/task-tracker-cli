package handlers

import (
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/args"
)

type CountHandler struct {
	from int64
	to   int64
}

func (h *CountHandler) SetFrom(from int64) {
	h.from = from
}

func (h *CountHandler) SetTo(to int64) {
	h.to = to
}

func (h *CountHandler) Run(params args.Params) error {
	if h.from > h.to && h.to > 0 {
		tmp := h.to
		h.to = h.from
		h.from = tmp
	}

	cnt := int64(len(params))
	if cnt < h.from && h.from > 0 && h.to <= 0 {
		return fmt.Errorf("invalid params count. It should be minimum %v", h.from)
	}
	if cnt > h.to && h.to > 0 && h.from <= 0 {
		return fmt.Errorf("invalid params count. It should be less than %v", h.to)
	}
	if cnt < h.from || cnt > h.to {
		if h.from == h.to {
			return fmt.Errorf("invalid params count. It should be %v", h.to)
		}
		if h.from > 0 && h.to > 0 {
			return fmt.Errorf("invalid params count. It should be between %v and %v", h.from, h.to)
		}
	}
	return nil
}
