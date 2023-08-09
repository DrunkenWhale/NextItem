package gen

import "math/big"

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Point struct {
	X int
	Y int
}

func LagrangianInterpolation(points []*Point) {
	
}

func calculateLagrangianCoefficient(index int, points []*Point) {
	xj := points[index].X
	res := big.NewInt(1)
	tmp := big.NewInt(0)
	for i := 0; i < len(points); i++ {
		res = tmp.Mul(big.NewInt(int64(xj-points[i].X)), res)
	}
}
