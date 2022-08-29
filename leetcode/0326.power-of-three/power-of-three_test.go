package problem0326

import (
	"fmt"
	"testing"
)

func Test_IsPowerOfThree(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println("i: ", i, "------------isPowerOfThree_1: ", isPowerOfThree_1(i))
		fmt.Println("i: ", i, "------------isPowerOfThree_2: ", isPowerOfThree_2(i))
	}
}