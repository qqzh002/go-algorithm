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

func (s *HeapSuite) SetUpSuite(c *C) {
	s.h = Heap([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
}

func (s *HeapSuite) TestParent(c *C) {
	var r int
	var e error

	r, e = s.h.Parent(2)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = s.h.Parent(3)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = s.h.Parent(4)
	c.Check(r, Equals, 2)
	c.Check(e, IsNil)

	r, e = s.h.Parent(9)
	c.Check(r, Equals, 4)
	c.Check(e, IsNil)

	r, e = s.h.Parent(10)
	c.Check(r, Equals, 5)
	c.Check(e, IsNil)

	_, e = s.h.Parent(1)
	c.Check(e, Equals, ErrNoParent)

	_, e = s.h.Parent(0)
	c.Check(e, Equals, ErrOutOfIndex)

	_, e = s.h.Parent(-1)
	c.Check(e, Equals, ErrOutOfIndex)

	_, e = s.h.Parent(11)
	c.Check(e, Equals, ErrOutOfIndex)
}

func (s *HeapSuite) TestLeft(c *C) {
	var r int
	var e error

	r, e = s.h.Left(1)
	c.Check(r, Equals, 2)
	c.Check(e, IsNil)

	r, e = s.h.Left(2)
	c.Check(r, Equals, 4)
	c.Check(e, IsNil)

	r, e = s.h.Left(5)
	c.Check(r, Equals, 10)
	c.Check(e, IsNil)

	_, e = s.h.Left(6)
	c.Check(e, Equals, ErrNoLeft)

	_, e = s.h.Left(10)
	c.Check(e, Equals, ErrNoLeft)

	_, e = s.h.Left(0)
	c.Check(e, Equals, ErrOutOfIndex)

	_, e = s.h.Left(-1)
	c.Check(e, Equals, ErrOutOfIndex)

	_, e = s.h.Left(11)
	c.Check(e, Equals, ErrOutOfIndex)
}
