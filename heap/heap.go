package heap

import "errors"

type Heap []int

var ErrNoParent = errors.New("no parent")

func Parent(i int) (int, error) {
	if i == 0 {
		return 0, ErrNoParent
	}
	return i >> 1, nil
}
