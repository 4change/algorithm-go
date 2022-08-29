package problem0342

import (
	"fmt"
	"testing"
)

func Test_IsPowerOfFour(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println("i: ", i, "------------isPowerOfFour: ", isPowerOfFour(i))
	}
}