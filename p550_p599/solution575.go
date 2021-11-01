package p550_p599

/*
575. 分糖果

给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。
你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

注意:
数组的长度为[2, 10,000]，并且确定为偶数。
数组中数字的大小在范围[-100,000, 100,000]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-candies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func distributeCandies(candyType []int) int {
	checkMap := make(map[int]bool, len(candyType))
	for _, candy := range candyType {
		checkMap[candy] = true
	}
	length := len(checkMap)
	if length <= len(candyType)/2 {
		return length
	}
	return len(candyType) / 2
}
