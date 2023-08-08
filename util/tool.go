package util

import "math/big"

func GCD(x, y int) int {
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

// RFCD
// 原谅我是真的不知道这个词怎么写 大概是分子分母通分 我参考了这个短语
// reduction of fractions to a common denominator
func RFCD(n, d *big.Int) (*big.Int, *big.Int) {
	gcd := big.NewInt(0).GCD(nil, nil, n, d)
	n = n.Div(n, gcd)
	d = d.Div(d, gcd)
	return n, d
}
