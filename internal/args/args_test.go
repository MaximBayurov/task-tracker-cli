package args

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type ParsedTestData struct {
	Raw    []string
	Parsed Arguments
}

func TestGetParsed(t *testing.T) {
	currArgs := os.Args
	for i, data := range getParsedTestData() {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			os.Args = data.Raw
			parsed := GetParsed()
			if !reflect.DeepEqual(parsed, data.Parsed) {
				t.Errorf("Unexpected parse results %v. %v expected for \"%v\"", parsed, data.Parsed, data.Raw)
			}
		})
	}
	os.Args = currArgs
}

func getParsedTestData() []ParsedTestData {
	return []ParsedTestData{
		{nil, Arguments{}},
		{[]string{"path", "cmd", "param1"}, Arguments{"cmd", Params{"param1"}}},
		{[]string{"path", "cmd"}, Arguments{"cmd", Params{}}},
		{[]string{"path", "cmd", "param1", "param1", "param1"}, Arguments{"cmd", Params{"param1", "param1", "param1"}}},
	}
}
