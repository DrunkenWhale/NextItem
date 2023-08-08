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
	big.Int{}
}

func gcd(x, y int) int {
	if x < y {
		x, y = swap(x, y)
	}
	return _gcd(x, y)
}

func _gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return _gcd(y, x%y)
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
