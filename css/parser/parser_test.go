package parser

import (
	"testing"
)

func Test_ParseSimple(t *testing.T) {
	err := Parse([]byte(`
foo > * .baz123
{
    color: red;
}
`))
	if err != nil {
		t.Error(err)
	}
}
