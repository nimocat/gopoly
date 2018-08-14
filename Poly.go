package gopoly

import (
	"fmt"
	"math/big"
)

func polyMul(p []*big.Int, value *big.Int) *big.Int {
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
