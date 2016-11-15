# go-iter

Basic iterator types for go.

## Usage

    npm install shadowmint/go-iterator --save

Then use `Next()` to iterate over values:

    i1 := iter.FromList(container)
    i2 := iter.FromList(container2)
    i3 := iter.Join(i1, i2)

		total := 0
		for val, err := i3.Next(); err == nil; val, err = i3.Next() {
			total += val.(int)
		}

By convention you can check if the iterator failed or hit end of iteration using:

    _, err := i3.Next()
		if errors.Is(err, iter.ErrEndIteration{}) {
      ...
    }

## Tests

    npm install
    npm run generate
    npm test
