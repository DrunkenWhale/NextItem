package gen

import "testing"

func TestLagrangianInterpolation(t *testing.T) {
	points := []*Point{
		{1, 1},
		{2, 3},
		{4, 5},
		{6, 7},
		{114, 514},
	}
	res := LagrangianInterpolation(points)
	t.Log(res.String())
}
