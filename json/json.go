// +build !jsoniterator

package json

import (
	"encoding/json"
)

var (
	// Marshal exports json.Marshal.
	Marshal = json.Marshal

	// Unmarshal exports json.Unmarshal.
	Unmarshal = json.Unmarshal

	// NewEncoder exports json.NewEncoder.
	NewEncoder = json.NewEncoder

	// NewDecoder exports json.NewDecoder.
	NewDecoder = json.NewDecoder
)
