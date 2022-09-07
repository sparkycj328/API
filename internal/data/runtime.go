package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type, which has the underlying type
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Wraps string in double quotes
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}