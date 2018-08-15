package gopoly

import (
	"fmt"
	"math/big"
)

type Poly []*big.Int

func PolyVal(p Poly, value *big.Int) *big.Int {
	length := len(p) - 1
	result := big.NewInt(0)
	j := 0
	for i := length; i >= 0; i-- {
		temp := big.NewInt(0)
		result.Add(result, temp.Mul(p[j], temp.Exp(value, big.NewInt(int64(i)), big.NewInt(0))))
		j++
		fmt.Println("No.", j, "calculate=", temp)
	}
	return result
}

func PolyMul(poly1 Poly, poly2 Poly) Poly {
	temp := new(big.Int)
	result := make([]*big.Int, len(poly1)+len(poly2)-1)
	for i := range result {
		result[i] = big.NewInt(0)
	} //initialize
	for i, coefficient := range poly1 {
		for j, coefficient2 := range poly2 {
			temp = temp.Mul(coefficient, coefficient2)
			result[i+j] = result[i+j].Add(result[i+j], temp)
		}
	}
	return result
}
