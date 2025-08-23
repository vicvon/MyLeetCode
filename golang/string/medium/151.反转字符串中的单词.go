/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30202
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	str := []byte(s)
	slow, fast := 0, 0
	// 删除头部指针
	for fast < len(str) && str[fast] == ' ' {
		fast++
	}
	// 删除单词中多余空格
	for fast < len(str) {
		if fast-1 >= 0 && str[fast] == ' ' && str[fast] == str[fast-1] {
			fast++
			continue
		}
		str[slow] = str[fast]
		fast++
		slow++
	}
	// 删除尾部空格
	if slow-1 >= 0 && str[slow-1] == ' ' {
		str = str[:slow-1]
	} else {
		str = str[:slow]
	}
	reverseString(str, 0, len(str)-1)
	start := 0
	end := 0
	for end < len(str) {
		if str[end] == ' ' {
			reverseString(str, start, end-1)
			start = end + 1
		}
		end++
	}

	reverseString(str, start, end-1)
	return string(str)
}

func reverseString(str []byte, start, end int) []byte {
	i, j := start, end
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return str
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/

