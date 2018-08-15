package gopoly

import (
	"fmt"
	"math/big"
	"testing"
)

var Poly1 []*big.Int = []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(2)}
var Poly2 []*big.Int = []*big.Int{big.NewInt(2), big.NewInt(4), big.NewInt(7)}

func TestPolyVal(t *testing.T) {
	if result := PolyVal(Poly1, big.NewInt(3)); result.Cmp(big.NewInt(44)) == 0 {
		t.Log("TestPolyVal successed")
	} else {
		t.Error("TestPolyVal failed")
	} //the result should be 2 + 5 * 3^1 + 3 * 3^2 = 44
}

func TestPolyMul(t *testing.T) {
	fmt.Println(PolyMul(Poly1, Poly2))
}
