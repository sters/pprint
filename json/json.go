package json

import (
	"encoding/json"
	"io"
)

// PPrint is aimed for JSON
type PPrint struct{}

func (*PPrint) GetType() string {
	return "json"
}

func (*PPrint) ParseAndPrint(r io.Reader, w io.Writer) error {
	var data interface{}
	if err := json.NewDecoder(r).Decode(data); err != nil {
		return err
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
}
