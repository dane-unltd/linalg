package linalg

import (
	"github.com/dane-unltd/linalg/lapacke"
	"github.com/gonum/blas/cblas"
	"github.com/gonum/matrix/mat64"
	check "launchpad.net/gocheck"
	"math"
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

func (s *S) TestSVD(c *check.C) {
	for _, t := range []struct {
		a *mat64.Dense

		epsilon float64
		small   float64

		wantu bool
		u     *mat64.Dense

		sigma []float64

		wantv bool
		v     *mat64.Dense
	}{
		{
			a: mustDense(mat64.NewDense(4, 2, []float64{2, 4, 1, 3, 0, 0, 0, 0})),

			epsilon: math.Pow(2, -52.0),
			small:   math.Pow(2, -966.0),

			wantu: true,
			u: mustDense(mat64.NewDense(4, 2, []float64{
				0.8174155604703632, -0.5760484367663209,
				0.5760484367663209, 0.8174155604703633,
				0, 0,
				0, 0,
			})),

			sigma: []float64{5.464985704219041, 0.365966190626258},

			wantv: true,
			v: mustDense(mat64.NewDense(2, 2, []float64{
				0.4045535848337571, -0.9145142956773044,
				0.9145142956773044, 0.4045535848337571,
			})),
		},
		{
			a: mustDense(mat64.NewDense(4, 2, []float64{2, 4, 1, 3, 0, 0, 0, 0})),

			epsilon: math.Pow(2, -52.0),
			small:   math.Pow(2, -966.0),

			wantu: true,
			u: mustDense(mat64.NewDense(4, 2, []float64{
				0.8174155604703632, -0.5760484367663209,
				0.5760484367663209, 0.8174155604703633,
				0, 0,
				0, 0,
			})),

			sigma: []float64{5.464985704219041, 0.365966190626258},

			wantv: false,
		},
		{
			a: mustDense(mat64.NewDense(4, 2, []float64{2, 4, 1, 3, 0, 0, 0, 0})),

			epsilon: math.Pow(2, -52.0),
			small:   math.Pow(2, -966.0),

			wantu: false,

			sigma: []float64{5.464985704219041, 0.365966190626258},

			wantv: true,
			v: mustDense(mat64.NewDense(2, 2, []float64{
				0.4045535848337571, -0.9145142956773044,
				0.9145142956773044, 0.4045535848337571,
			})),
		},
		{
			a: mustDense(mat64.NewDense(4, 2, []float64{2, 4, 1, 3, 0, 0, 0, 0})),

			epsilon: math.Pow(2, -52.0),
			small:   math.Pow(2, -966.0),

			sigma: []float64{5.464985704219041, 0.365966190626258},
		},
		{
			// FIXME(kortschak)
			// This test will fail if t.sigma is set to the real expected values
			// or if u and v are requested, due to a bug in the original Jama code
			// forcing a to be a tall or square matrix.
			//
			// This is a failing case to use to fix that bug.
			a: mustDense(mat64.NewDense(3, 11, []float64{
				1, 1, 0, 1, 0, 0, 0, 0, 0, 11, 1,
				1, 0, 0, 0, 0, 0, 1, 0, 0, 12, 2,
				1, 1, 0, 0, 0, 0, 0, 0, 1, 13, 3,
			})),

			epsilon: math.Pow(2, -52.0),
			small:   math.Pow(2, -966.0),

			// FIXME(kortschak) sigma is one element longer than it should be.
			sigma: []float64{21.25950088109745, 1.5415021616856577, 1.2873979074613637, 0},
		},
	} {
		svd := lapacke.SVD(mat64.DenseCopyOf(t.a), t.epsilon, t.small, t.wantu, t.wantv)
		if t.sigma != nil {
			c.Check(svd.Sigma, check.DeepEquals, t.sigma)
		}
		s := svd.S()

		if svd.U != nil {
			c.Check(svd.U.Equals(t.u), check.Equals, true)
		} else {
			c.Check(t.wantu, check.Equals, false)
			c.Check(t.u, check.IsNil)
		}
		if svd.V != nil {
			c.Check(svd.V.Equals(t.v), check.Equals, true)
		} else {
			c.Check(t.wantv, check.Equals, false)
			c.Check(t.v, check.IsNil)
		}

		if t.wantu && t.wantv {
			c.Assert(svd.U, check.NotNil)
			c.Assert(svd.V, check.NotNil)
			vt := &mat64.Dense{}
			vt.TCopy(svd.V)
			svd.U.Mul(svd.U, s)
			svd.U.Mul(svd.U, vt)
			c.Check(svd.U.EqualsApprox(t.a, 1e-12), check.Equals, true)
		}
	}
}

func (s *S) TestCholesky(c *check.C) {
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
		c.Check(lc.EqualsApprox(t.a, 1e-4), check.Equals, true)

		x := cf.Solve(eye())

		t.a.Mul(t.a, x)
		c.Check(t.a.EqualsApprox(eye(), 1e-4), check.Equals, true)
	}
}
