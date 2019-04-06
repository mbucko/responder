package processor

import (
	"errors"
	"strings"
	"testing"
)

func TestProcessString(t *testing.T) {

	type Ret struct {
		ret string
		err error
	}

	cases := []struct {
		in  string
		ret Ret
	}{
		{"test", Ret{"Upper case of 'test' is 'TEST'.", nil}},
		{"", Ret{"", errors.New("Input string is empty")}},
	}

	for _, cases := range cases {
		// process(cases.in)
		ret, err := processString(cases.in)
		if len(ret) == 0 && len(cases.ret.ret) != 0 {
			t.Errorf("Test failed, the expected return value should be empty but it's not, value = %s", cases.ret.ret)
		} else if len(ret) != 0 && ret == cases.ret.ret {
			t.Errorf("Test failed, the expected return value '%s' is not '%s'", cases.ret.ret, ret)
		}
		if cases.ret.err == nil {
			if err != nil {
				t.Errorf("Test failed, the expected nil error value not nil, error = '%v'", err)
			}
		} else if cases.ret.err.Error() != err.Error() {
			t.Errorf("Test failed, the expected error value '%v' is not '%v'", cases.ret.err, err)
		}
	}
}

func TestProcess(t *testing.T) {
	reader := strings.NewReader("{ \"id\": 0, \"toCaps\" : \"nasa\" }")
	output := process(reader)
	const expected = "{\"id\":0,\"toCaps\":\"NASA\"}"
	if output != expected {
		t.Errorf("Test failed, the expected error value '%s' is not '%s'", expected, output)
	}
}
