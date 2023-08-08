package main

import (
	"fmt"
	"math/big"
)

func main() {
	//seq := []int{1, 1, 4, 5, 1, 4, 1919810}
	a := big.NewInt(2)
	a.Mul(big.NewInt(111), big.NewInt(11))
	fmt.Println(a)
	a.Mul(big.NewInt(111), big.NewInt(11))
	fmt.Println(a)
}
