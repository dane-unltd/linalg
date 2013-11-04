package linalg

import (
	"github.com/dane-unltd/linalg/lapacke"
	"github.com/gonum/blas/cblas"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/matrix/mat64/la"
	check "launchpad.net/gocheck"
	"testing"
)

// Helpers

func mustDense(m *mat64.Dense, err error) *mat64.Dense {
	if err != nil {
		panic(err)
	}
	return m
}

func eye() *mat64.Dense {
	return mustDense(mat64.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}))
}

// Tests

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

func (s *S) SetUpSuite(c *check.C) { mat64.Register(cblas.Blas{}) }

var _ = check.Suite(&S{})

func (s *S) TestCholeskyLA(c *check.C) {
	for _, t := range []struct {
		a   *mat64.Dense
		spd bool
	}{
		{
			a: mustDense(mat64.NewDense(3, 3, []float64{
				4, 1, 1,
				1, 2, 3,
				1, 3, 6,
			})),

			spd: true,
		},
	} {
		cf := la.Cholesky(t.a)
		c.Check(cf.SPD, check.Equals, t.spd)

		lt := &mat64.Dense{}
		lt.TCopy(cf.L)
		lc := mat64.DenseCopyOf(cf.L)

		lc.Mul(lc, lt)
		c.Check(lc.EqualsApprox(t.a, 1e-12), check.Equals, true)

		x := cf.Solve(eye())

		t.a.Mul(t.a, x)
		c.Check(t.a.EqualsApprox(eye(), 1e-12), check.Equals, true)
	}
}

func (s *S) TestCholeskyLAPACKE(c *check.C) {
	for _, t := range []struct {
		a   *mat64.Dense
		spd bool
	}{
		{
			a: mustDense(mat64.NewDense(3, 3, []float64{
				4, 1, 1,
				1, 2, 3,
				1, 3, 6,
			})),

			spd: true,
		},
	} {
		cf := lapacke.Cholesky(t.a)
		c.Check(cf.SPD, check.Equals, t.spd)

		lt := &mat64.Dense{}
		lt.TCopy(cf.L)
		lc := mat64.DenseCopyOf(cf.L)

		lc.Mul(lc, lt)
		c.Check(lc.EqualsApprox(t.a, 1e-12), check.Equals, true)

		x := cf.Solve(eye())

		t.a.Mul(t.a, x)
		c.Check(t.a.EqualsApprox(eye(), 1e-12), check.Equals, true)
	}
}
