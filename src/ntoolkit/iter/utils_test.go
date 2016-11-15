package iter_test

import (
	"container/list"
	"ntoolkit/assert"
	"ntoolkit/errors"
	"ntoolkit/iter"
	"testing"
)

func TestCount(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		container := list.New()
		container.PushBack(100)
		container.PushBack(110)
		container.PushBack(111)

		i := iter.FromList(container)
		count, err := iter.Count(i)
		T.Assert(err == nil)
		T.Assert(count == 3)

		_, err = iter.Count(i)
		T.Assert(errors.Is(err, iter.ErrEndIteration{}))
	})
}

func TestCollect(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		container := list.New()
		container.PushBack(100)
		container.PushBack(110)
		container.PushBack(111)

		i := iter.FromList(container)
		all, err := iter.Collect(i)
		T.Assert(err == nil)
		T.Assert(all != nil)
		T.Assert(len(all) == 3)

		_, err = iter.Count(i)
		T.Assert(errors.Is(err, iter.ErrEndIteration{}))
	})
}
