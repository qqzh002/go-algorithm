package heap

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestParent(c *C) {
	var r int
	var e error

	r, e = Parent(1)
	c.Check(r, Equals, 0)
	c.Check(e, IsNil)

	r, e = Parent(2)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = Parent(3)
	c.Check(r, Equals, 1)
	c.Check(e, IsNil)

	r, e = Parent(4)
	c.Check(r, Equals, 2)
	c.Check(e, IsNil)

	_, e = Parent(0)
	c.Check(e, Equals, ErrNoParent)
}
