package json

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPPrint_ParseAndPrint(t *testing.T) {
	tests := map[string]struct {
		input      string
		wantOutput string
		wantError  bool
	}{
		"invalid": {
			`{"aaa":1`,
			``,
			true,
		},
		"object": {
			`{"aaa":{"bbb":"ccc"}}`,
			`{
    "aaa": {
        "bbb": "ccc"
    }
}
`,
			false,
		},
		"array": {
			`[1,2,3,[4,5]]`,
			`[
    1,
    2,
    3,
    [
        4,
        5
    ]
]
`,
			false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			input := bytes.NewBuffer([]byte(test.input))
			output := bytes.NewBuffer([]byte{})
			err := (&PPrint{}).ParseAndPrint(input, output)

			if test.wantError {
				assert.Error(t, err)
				return // test end
			}

			assert.NoError(t, err)

			o := string(output.Bytes())
			assert.Equal(t, test.wantOutput, o)
		})
	}
}
