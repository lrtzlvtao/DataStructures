package translateNum

import "strconv"

/*
给定一个数字，我们按照如下规则把它翻译为字符串：
	0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：
0 <= num < 2^31
*/

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
		
	tmp := num % 100
	if tmp < 10 || tmp > 25 {
		return translateNum(num / 10)
	}
	return translateNum(num / 10) + translateNum(num / 100)
}

func translateNumDp(num int) int {
    src := strconv.Itoa(num)
    p, q, r := 0, 0, 1
    for i := 0; i < len(src); i++ {
        p, q, r = q, r, 0
        r += q
        if i == 0 {
            continue
        }
        pre := src[i-1:i+1]
        if pre <= "25" && pre >= "10" {
            r += p
        }
    }
    return r
}
