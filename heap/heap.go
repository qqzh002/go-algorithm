package heap

import "errors"

type Heap []int

var ErrOutOfIndex = errors.New("out of index")
var ErrNoParent = errors.New("no parent")

func (h *Heap) Parent(i int) (int, error) {
	if i <= 0 || i > len(*h) {
		return 0, ErrOutOfIndex
	}
	if i == 1 {
		return 0, ErrNoParent
	}
	return i >> 1, nil
}

func (h *Heap) Left(i int) (int, error) {
	return i << 1, nil
}
