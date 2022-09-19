package main

func isValid(s string) bool {
    pair := map[byte]byte {
        ')':'(',
        ']':'[',
        '}':'{',
    }

    stack := []byte{}

// 左括号入栈，右括号与栈顶元素匹配， 成功 ==> 弹出， 失败 ==> 返回false
// 从左到右扫描结束，检查栈是否为空， 空 ==> true, 不为空 ==> false
    for i := 0; i < len(s); i++ {
        if _, ok := pair[s[i]]; ok {	// 右括号与栈顶元素匹配
            // 匹配过程
            if len(stack) == 0 || stack[len(stack) - 1] != pair[s[i]] {		// 失败，返回false
                return false
            }
            stack = stack[:len(stack) - 1]	// 成功，弹出
        } else {	// 左括号入栈
            stack = append(stack, s[i])
        }
    }

    return len(stack) == 0
}