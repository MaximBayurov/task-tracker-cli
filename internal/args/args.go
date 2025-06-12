package args

import (
	"os"
)

// Params data type describing program's params
type Params []string

// Arguments contains command code and params that will pass in it
type Arguments struct {
	Cmd    string
	Params Params
}

// GetParsed return parsed cli params
func GetParsed() Arguments {
	var args Arguments

	if len(os.Args) < 1 {
		return args
	}
	raw := os.Args[1:]
	if len(raw) < 1 {
		return args
	}

	if raw[0] != "" {
		args.Cmd = raw[0]
	}

	args.Params = Params{}
	if len(raw) > 1 {
		args.Params = raw[1:]
	}

	return args
}
