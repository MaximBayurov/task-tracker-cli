package mark

import "fmt"

// Mark the same as status
type Mark int8

const (
	Todo       Mark = 10
	InProgress Mark = 20
	Done       Mark = 30
	Canceled   Mark = 40
)

// FromString return Mark from passed string
func FromString(string string) (Mark, error) {
	var result Mark
	result, ok := AllowedMarks[string]
	if !ok {
		return 0, fmt.Errorf("unsupported mark \"%v\"", string)
	}
	return result, nil
}

// ToString return string from passed Mark
func (m Mark) ToString() string {
	var result string
	for k, am := range AllowedMarks {
		if am == m {
			result = k
			break
		}
		if result == "" {
			result = k
		}
	}
	return result
}

var AllowedMarks = map[string]Mark{
	"todo":        Todo,
	"in-progress": InProgress,
	"canceled":    Done,
	"done":        Canceled,
}
