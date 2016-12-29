package iter

import "ntoolkit/errors"

// Count enumerates an iterator, consuming it and returning the length.
func Count(iterator Iter) (int, error) {
	count := 0
	var err error = nil
	for _, err = iterator.Next(); err == nil; _, err = iterator.Next() {
		count++
	}
	if !errors.Is(err, ErrEndIteration{}) {
		return 0, err
	}
	return count, nil
}

// Collect enumerates an iterator, consuming it and returning a slice of the values.
func Collect(iterator Iter) ([]interface{}, error) {
	values := make([]interface{}, 0)
	var value interface{} = nil
	var err error = nil
	for value, err = iterator.Next(); err == nil; value, err = iterator.Next() {
		values = append(values, value)
	}
	if !errors.Is(err, ErrEndIteration{}) {
		return nil, err
	}
	return values, nil
}
