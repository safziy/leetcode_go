package p450_p499

import "unicode"

/*
482. 密钥格式化
有一个密钥字符串 S ，只包含字母，数字以及 '-'（破折号）。其中， N 个 '-' 将字符串分成了 N+1 组。

给你一个数字 K，请你重新格式化字符串，使每个分组恰好包含 K 个字符。特别地，第一个分组包含的字符个数必须小于等于 K，但至少要包含 1 个字符。两个分组之间需要用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。

给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。

提示:

S 的长度可能很长，请按需分配大小。K 为正整数。
S 只包含字母数字（a-z，A-Z，0-9）以及破折号'-'
S 非空
*/

func licenseKeyFormatting(s string, k int) string {
	cnt := 0
	for _, c := range s {
		if c != '-' {
			cnt++
		}
	}
	if cnt == 0 {
		return ""
	}
	var bytes []byte
	mod := cnt % k
	if mod != 0 {
		mod = k - mod
	}
	for _, c := range s {
		if c != '-' {
			bytes = append(bytes, byte(unicode.ToUpper(c)))
			mod++
			if mod == k {
				bytes = append(bytes, '-')
				mod = 0
			}
		}
	}
	return string(bytes[:len(bytes)-1])
}
