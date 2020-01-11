package heap

import "errors"

type Heap []int

var ErrOutOfIndex = errors.New("out of index")
var ErrNoParent = errors.New("no parent")
var ErrNoLeft = errors.New("no left")
var ErrNoRight = errors.New("no right")

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
	if i <= 0 || i > len(*h) {
		return 0, ErrOutOfIndex
	}
	r := i << 1
	if r > len(*h) {
		return 0, ErrNoLeft
	}
	return r, nil
}

func (h *Heap) Right(i int) (int, error) {
	if i <= 0 || i > len(*h) {
		return 0, ErrOutOfIndex
	}
	r := i<<1 + 1
	if r > len(*h) {
		return 0, ErrNoRight
	}
	return r, nil
}

func (h *Heap) get(i int) int {
	return (*h)[i-1]
}

func (h *Heap) exchange(i, j int) {
	(*h)[i-1], (*h)[j-1] = (*h)[j-1], (*h)[i-1]
}

func (h *Heap) MaxHeapify(i int) error {
	if i <= 0 || i > len(*h) {
		return ErrOutOfIndex
	}

	largest := i

	l, errLeft := h.Left(i)
	if errLeft == nil && h.get(l) > h.get(largest) {
		largest = l
	}

	r, errRight := h.Right(i)
	if errRight == nil && h.get(r) > h.get(largest) {
		largest = r
	}

	if largest != i {
		h.exchange(i, largest)
		h.MaxHeapify(largest)
	}

	return nil
}

func (h *Heap) BuildMaxHeap() {
	i, _ := h.Parent(len(*h))
	for i >= 1 {
		h.MaxHeapify(i)
		i--
	}
}
