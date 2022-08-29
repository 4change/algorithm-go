package problem0231

import (
	"fmt"
	"testing"
)

func Test_IsPowerOfFour(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println("i: ", i, "------------isPowerOfFour: ", isPowerOfTwo(i))
	}
}