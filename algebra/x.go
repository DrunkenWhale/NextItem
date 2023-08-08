package algebra

import (
	"Sequence/util"
	"fmt"
	"math/big"
)

type X struct {
	CoefficientNumerator   *big.Int
	CoefficientDenominator *big.Int
	Power                  int
}

func NewAlgebra(coefficientNumerator, coefficientDenominator *big.Int, power int) *X {
	return &X{
		CoefficientNumerator:   coefficientNumerator,
		CoefficientDenominator: coefficientDenominator,
		Power:                  power,
	}
}

// Add 返回二者相加的值 不会对原对象有所影响
func (x *X) Add(y *X) *X {
	if x.Power != y.Power {
		panic("power unequal")
	}
	on := big.NewInt(0).Add(
		big.NewInt(0).Mul(x.CoefficientNumerator, y.CoefficientDenominator),
		big.NewInt(0).Mul(y.CoefficientNumerator, x.CoefficientDenominator),
	)
	od := big.NewInt(0).Mul(x.CoefficientDenominator, y.CoefficientDenominator)
	nn, nd := util.RFCD(on, od)
	return NewAlgebra(nn, nd, x.Power)
}

// Mul 返回二者相乘的值 不会对原对象有所影响
func (x *X) Mul(y *X) *X {
	on := big.NewInt(0).Mul(x.CoefficientNumerator, y.CoefficientNumerator)
	od := big.NewInt(0).Mul(x.CoefficientDenominator, y.CoefficientDenominator)
	nn, nd := util.RFCD(on, od)
	return NewAlgebra(nn, nd, x.Power*y.Power)
}

// Sub 返回二者相减法的值 不会对原对象有所影响
func (x *X) Sub(y *X) *X {
	if x.Power != y.Power {
		panic("power unequal")
	}
	on := big.NewInt(0).Sub(
		big.NewInt(0).Mul(x.CoefficientNumerator, y.CoefficientDenominator),
		big.NewInt(0).Mul(y.CoefficientNumerator, x.CoefficientDenominator),
	)
	od := big.NewInt(0).Mul(x.CoefficientDenominator, y.CoefficientDenominator)
	nn, nd := util.RFCD(on, od)
	return NewAlgebra(nn, nd, x.Power)
}

// Div 返回二者相除的值 不会对原对象有所影响
func (x *X) Div(y *X) *X {
	on := big.NewInt(0).Mul(x.CoefficientNumerator, y.CoefficientDenominator)
	od := big.NewInt(0).Mul(x.CoefficientDenominator, y.CoefficientNumerator)
	nn, nd := util.RFCD(on, od)
	return NewAlgebra(nn, nd, x.Power-y.Power)
}

func (x *X) String() string {
	if x.Power == 0 {
		if x.CoefficientNumerator.Cmp(big.NewInt(0)) == 0 {
			return fmt.Sprintf("")
		}
		if x.CoefficientDenominator.Cmp(big.NewInt(1)) == 0 {
			if x.CoefficientNumerator.Cmp(big.NewInt(1)) == 0 {
				return fmt.Sprintf("")
			}
			return fmt.Sprintf("%v", x.CoefficientNumerator)
		}
		return fmt.Sprintf("\\frac{%v}{%v} ", x.CoefficientNumerator, x.CoefficientDenominator)
	}
	if x.CoefficientNumerator.Cmp(big.NewInt(0)) == 0 { // 分子为0
		return fmt.Sprintf("")
	}
	if x.CoefficientDenominator.Cmp(big.NewInt(1)) == 0 {
		if x.CoefficientNumerator.Cmp(big.NewInt(1)) == 0 {
			return fmt.Sprintf("x^{%v}", x.Power)
		}
		return fmt.Sprintf("%v x^{%v}", x.CoefficientNumerator, x.Power)
	}
	return fmt.Sprintf("\\frac{%v}{%v} x^{%v}", x.CoefficientNumerator, x.CoefficientDenominator, x.Power)
}
