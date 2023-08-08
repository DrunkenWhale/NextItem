package algebra

import (
	"math/big"
	"testing"
)

func TestAlgebra(t *testing.T) {
	x := NewAlgebra(big.NewInt(3), big.NewInt(4), 3)
	y := NewAlgebra(big.NewInt(3), big.NewInt(4), 3)
	t.Log(x.Add(y).String())
	t.Log(x.Mul(y).String())
	t.Log(x.Sub(y).String())
	t.Log(x.Div(y).String())
	t.Log(NewAlgebra(big.NewInt(7), big.NewInt(8), 0).String())
}
