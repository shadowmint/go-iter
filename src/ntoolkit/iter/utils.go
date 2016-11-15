package iter

import "ntoolkit/errors"

// Count enumerates an iterator, consuming it and returning the length.
func Count(iterator Iter) (int, error) {
	count := 0
	_, err := iterator.Next()
	if err != nil {
		return 0, err
	}
	for ; err == nil; _, err = iterator.Next() {
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
	val, err := iterator.Next()
	if err != nil {
		return nil, err
	}
	for ; err == nil; val, err = iterator.Next() {
		values = append(values, val)
	}
	if !errors.Is(err, ErrEndIteration{}) {
		return nil, err
	}
	return values, nil
}
