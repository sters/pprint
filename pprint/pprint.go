package pprint

import "io"

// PPrint interface
type PPrint interface {
	// GetType returns aimed what type for
	GetType() string

	// ParseAndPrint input text
	ParseAndPrint(r io.Reader, w io.Writer) error
}
