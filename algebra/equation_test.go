package algebra

import (
	"math/big"
	"testing"
)

func TestNewEquation(t *testing.T) {
	type test struct {
		a int
		b string
		c bool
	}
	a := &test{
		a: 1,
		b: "",
		c: true,
	}
	b := *a

	if a == &b {
		t.Fatal("")
	}
	if a.b != b.b {
		t.Fatal("")
	}
	a.b = "114"
	if a.b == b.b {
		t.Fatal("")
	}
}

func TestEquation_Operator(t *testing.T) {
	// x^2-7
	eqa := NewEquation([]*X{NewAlgebra(big.NewInt(1), big.NewInt(1), 2)}).Sub(NewAlgebra(big.NewInt(7), big.NewInt(1), 0))
	t.Log(eqa.Sort().String())
	// (x^2-7)*6x/7
	eqa = eqa.Mul(NewAlgebra(big.NewInt(6), big.NewInt(7), 1))
	t.Log(eqa.Sort().String())
	// (x^2-7)*6x/7*7*(4x^3-8x^7+51)
	eqa = eqa.Mul(NewAlgebra(big.NewInt(7), big.NewInt(1), 0)).MulEquation(NewEquation([]*X{
		{big.NewInt(4), big.NewInt(1), 3},
		{big.NewInt(-8), big.NewInt(1), 7},
		{big.NewInt(51), big.NewInt(1), 0},
	}))
	t.Log(eqa.Sort().String())
	// (x^2-7)*6x/7*7*(4x^3-8x^7+51)/114
	eqa = eqa.Div(NewAlgebra(big.NewInt(114), big.NewInt(1), 0))
	t.Log(eqa.Sort().String())
	// (x^2-7)*6x/7*7*(4x^3-8x^7+51)/114+(500x^17-8)
	eqa = eqa.AddEquation(NewEquation([]*X{
		NewAlgebra(big.NewInt(500), big.NewInt(1), 17),
		NewAlgebra(big.NewInt(-8), big.NewInt(1), 0),
	}))
	t.Log(eqa.Sort().String())
	// (x^2-7)*6x/7*7*(4x^3-8x^7+51)/114+(500x^17-8)+14/79
	eqa = eqa.Add(NewAlgebra(big.NewInt(14), big.NewInt(79), 0))
	t.Log(eqa.Sort().String())
}
