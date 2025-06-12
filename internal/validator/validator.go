package validator

import "github.com/MaximBayurov/task-tracker-cli/internal/args"

// Validator validator struct
type Validator struct {
	params   args.Params
	handlers []Handler
}

// AddHandler add handler to validator
func (v *Validator) AddHandler(handler Handler) {
	v.handlers = append(v.handlers, handler)
}

// SetParams sets params
func (v *Validator) SetParams(params args.Params) {
	v.params = params
}

// Run runs all validator's handlers
func (v *Validator) Run() []error {
	errs := make([]error, 0)
	for _, h := range v.handlers {
		err := h.Run(v.params)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
