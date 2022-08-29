package problem0231

func isPowerOfTwo(n int) bool {
    for n >= 2 && n%2 == 0 {
        n /= 2
    }

    return n == 1
}