package heap

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type HeapSuite struct {
	h Heap
}

var _ = Suite(&HeapSuite{})

func (s *HeapSuite) TestParent(c *C) {
	var r int
	var e error

	h := Heap([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})

	r, e = h.Parent(1)
	c.Check(r, Equals, 0)
	c.Check(e, IsNil)

	r, e = h.Parent(2)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = h.Parent(3)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = h.Parent(4)
	c.Check(r, Equals, 2)
	c.Check(e, IsNil)

	_, e = h.Parent(0)
	c.Check(e, Equals, ErrNoParent)

	_, e = h.Parent(-1)
	c.Check(e, Equals, ErrOutOfIndex)

	_, e = h.Parent(10)
	c.Check(e, Equals, ErrOutOfIndex)
}
