package main

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type TinyERPSuite struct {
	tiny *TinyERP
}

var _ = check.Suite(&TinyERPSuite{})

func (s *TinyERPSuite) SetUpSuite(c *check.C) {
	s.tiny = &TinyERP{}
}

func (s *TinyERPSuite) TestWorks(c *check.C) {
	c.Assert(s.tiny.works(), check.Equals, true)
}
