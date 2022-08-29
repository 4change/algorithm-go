package problem0326

func isPowerOfThree_1(n int) bool {
    for n > 0 && n%3 == 0 {
        n /= 3
    }
    return n == 1
}

func isPowerOfThree_2(n int) bool {
    return n > 0 && 1162261467%n == 0
}