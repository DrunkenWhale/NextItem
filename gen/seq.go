package gen

import (
	"Sequence/algebra"
	"math/big"
)

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

func LagrangianInterpolation(points []*Point) *algebra.Equation {
	res := algebra.NewEquation([]*algebra.X{})
	for i := range points {
		res = res.AddEquation(calculateLagrangianSingle(i, points))
	}
	return res.Sort()
}

// 一个拉格朗日插值法中的单项 不知道怎么拼写捏
func calculateLagrangianSingle(index int, points []*Point) *algebra.Equation {
	xj := points[index].X
	num := algebra.NewEquation([]*algebra.X{ // 分子
		{big.NewInt(1), big.NewInt(1), 0},
	})
	den := big.NewInt(1) // 分母
	for i := 0; i < len(points); i++ {
		if i == index {
			continue
		}
		eqa := algebra.NewEquation([]*algebra.X{
			{big.NewInt(1), big.NewInt(1), 1},
		})
		num = num.MulEquation(eqa.Sub(algebra.NewAlgebra(big.NewInt(int64(points[i].X)), big.NewInt(1), 0)))
		den = big.NewInt(0).Mul(big.NewInt(int64(xj-points[i].X)), den)
	}
	return num.Div(algebra.NewAlgebra(den, big.NewInt(1), 0)).
		Mul(algebra.NewAlgebra(big.NewInt(int64(points[index].Y)), big.NewInt(1), 0))
}
