package algebra

import (
	"fmt"
	"math/big"
	"sort"
)

type Equation struct {
	arr []*X
}

func NewEquation(arr []*X) *Equation {
	eqa := new(Equation)
	for i := 0; i < len(arr); i++ {
		e := *arr[i] //拷贝一份结构体 避免因为指针的问题出现预期之外的改动
		eqa.arr = append(eqa.arr, &e)
	}
	return eqa
}

func (eq *Equation) Add(x *X) (eqa *Equation) {
	eqa = NewEquation(eq.arr)
	for i := 0; i < len(eqa.arr); i++ {
		if eqa.arr[i].Power == x.Power {
			eqa.arr[i] = eqa.arr[i].Add(x)
			return
		}
	}
	// 到这了说明没有同次项的
	eqa.arr = append(eqa.arr, x)
	return
}

func (eq *Equation) AddEquation(other *Equation) (eqa *Equation) {
	eqa = NewEquation(eq.arr)
	for _, x := range other.arr { //O(n^2)
		eqa = eqa.Add(x) // 开销爆表照样用好吧
	}
	return
}

func (eq *Equation) Sub(x *X) (eqa *Equation) {
	eqa = NewEquation(eq.arr)
	for i := 0; i < len(eqa.arr); i++ {
		if eqa.arr[i].Power == x.Power {
			eqa.arr[i] = eqa.arr[i].Sub(x)
			return
		}
	}
	// 取负值
	eqa.arr = append(eqa.arr, NewAlgebra(big.NewInt(0), big.NewInt(1), x.Power).Sub(x))
	return
}

func (eq *Equation) SubEquation(other *Equation) (eqa *Equation) {
	eqa = NewEquation(eq.arr)
	for _, x := range other.arr {
		eqa = eqa.Sub(x)
	}
	return
}

func (eq *Equation) Mul(x *X) (eqa *Equation) {
	eqa = NewEquation(eq.arr)
	for i := 0; i < len(eqa.arr); i++ {
		eqa.arr[i] = eqa.arr[i].Mul(x)
	}
	return
}

func (eq *Equation) MulEquation(other *Equation) (eqa *Equation) {
	m := make(map[int][]*X) // power => X
	for _, x := range other.arr {
		newEqa := eq.Mul(x)
		for _, xx := range newEqa.arr {
			m[xx.Power] = append(m[xx.Power], xx)
		}
	}
	res := make([]*X, 0)
	for _, arr := range m {
		sum := arr[0]
		for i := 1; i < len(arr); i++ {
			sum = sum.Add(arr[i])
		}
		res = append(res, sum)
	}
	eqa = NewEquation(res)
	return
}

func (eq *Equation) Div(x *X) (eqa *Equation) {
	// 那话怎么说来着 ÷一个数等于×这个数的倒数 小学数学赛高!
	reverseX := NewAlgebra(x.CoefficientDenominator, x.CoefficientNumerator, -x.Power)
	return eq.Mul(reverseX)
}

func (eq *Equation) DivEquation(other *Equation) (eqa *Equation) {
	panic("我想到了一种美妙的写法 可是这行空白太小写不下(其实就是不会)")
}

// Sort 让等式变的有序
// 顺带一提这也是不会影响原对象的
func (eq *Equation) Sort() *Equation {
	eqa := NewEquation(eq.arr)
	sort.Slice(eqa.arr, func(i, j int) bool {
		return eqa.arr[i].Power > eqa.arr[j].Power
	})
	return eqa
}

func (eq *Equation) String() string {
	if len(eq.arr) == 0 {
		return ""
	}
	str := ""
	zero := NewAlgebra(big.NewInt(0), big.NewInt(1), 0)
	for i := 0; i < len(eq.arr); i++ {
		cmpRes := eq.arr[i].CmpCoefficient(zero)
		if cmpRes == 0 { // 系数为0 不会打印出来
			continue
		}
		if cmpRes > 0 {
			str += fmt.Sprintf(" + %v", eq.arr[i].String())
		} else {
			// 为负数的话 就乘一个负数然后手动打印减号呗
			str += fmt.Sprintf(" - %v", eq.arr[i].Mul(NewAlgebra(big.NewInt(-1), big.NewInt(1), 0)).String())
		}
	}
	if str[:3] == " + " {
		str = str[3:]
	}
	return str
}
