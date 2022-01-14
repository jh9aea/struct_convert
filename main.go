package struct_convert

import (
	"bytes"
	"encoding/gob"
)

// converts one struct to another through bytes buffer
func Convert[To, From any](from From) (To, error) {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(&from)
	var to To
	if err != nil {
		return to, err
	}
	err = gob.NewDecoder(&buf).Decode(&to)
	return to, err
}
