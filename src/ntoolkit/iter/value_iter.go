package iter

import (
	"ntoolkit/errors"
)

// ValueIter implements Iterator for an arbitrary single value
type ValueIter struct {
	value interface{}
	end   bool
}

// FromList returns a new list iterator for a list
func FromValue(value interface{}) *ValueIter {
	return &ValueIter{value, false}
}

// Return the single item or raise an error
func (iterator *ValueIter) Next() (interface{}, error) {
	if iterator.end {
		return nil, errors.Fail(ErrEndIteration{}, nil, "No more values")
	}
	iterator.end = true
	return iterator.value, nil
}
