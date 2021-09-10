package util

import (
	"fmt"
	"math/big"
	"testing"
)

func TestQuickSortBigInt(t *testing.T) {
	arr := make([]*big.Int, 100)
	for i := len(arr) - 1; i >=0 ; i-- {
		arr[i] = big.NewInt(int64(i))
	}

	QuickSortBigInt(arr,0,len(arr)-1)

	for i, b := range arr {
		fmt.Println(i,"==>",b.String())
	}
}
