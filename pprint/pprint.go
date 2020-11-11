package pprint

import (
	"fmt"
	"io"
	"strings"

	"github.com/sters/pprint/json"
)

// PPrint interface
type PPrint interface {
	// GetType returns aimed what type for
	GetType() string

	// ParseAndPrint input text
	ParseAndPrint(r io.Reader, w io.Writer) error
}

// Nop type as error
type Nop struct {
	types []string
}

func (*Nop) GetType() string { return "nop" }
func (n *Nop) ParseAndPrint(r io.Reader, w io.Writer) error {
	return fmt.Errorf(
		"need to choose type: %v",
		strings.Join(n.types, ","),
	)
}

func ChooseType(t string) PPrint {
	types := []PPrint{
		&json.PPrint{},
	}

	typeList := make([]string, 0, len(types))
	for _, tt := range types {
		if tt.GetType() == t {
			return tt
		}
		typeList = append(typeList, tt.GetType())
	}

	return &Nop{typeList}
}
