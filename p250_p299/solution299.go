package p250_p299

import "fmt"

/*
你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字secret 和朋友猜测的数字guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。


提示：

1 <= secret.length, guess.length <= 1000
secret.length == guess.length
secret 和 guess 仅由数字组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bulls-and-cows
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getHint(secret string, guess string) string {
	length := len(secret)
	aCount, bCount := 0, 0
	secretMap := map[uint8]int{}
	guessMap := map[uint8]int{}
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			aCount++
		} else {
			secretMap[secret[i]]++
			guessMap[guess[i]]++
		}
	}
	for num, count := range secretMap {
		bCount += min(count, guessMap[num])
	}
	return fmt.Sprintf("%dA%dB", aCount, bCount)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
