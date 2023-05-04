package p1150_p1199

import "sort"

//我们把无限数量 ∞ 的栈排成一行，按从左到右的次序从 0 开始编号。每个栈的的最大容量 capacity 都相同。
//
// 实现一个叫「餐盘」的类 DinnerPlates：
//
//
// DinnerPlates(int capacity) - 给出栈的最大容量 capacity。
// void push(int val) - 将给出的正整数 val 推入 从左往右第一个 没有满的栈。
// int pop() - 返回 从右往左第一个 非空栈顶部的值，并将其从栈中删除；如果所有的栈都是空的，请返回 -1。
// int popAtStack(int index) - 返回编号 index 的栈顶部的值，并将其从栈中删除；如果编号 index 的栈是空的，请返回 -
//1。
//
//
//
//
// 示例：
//
// 输入：
//["DinnerPlates","push","push","push","push","push","popAtStack","push","push",
//"popAtStack","popAtStack","pop","pop","pop","pop","pop"]
//[[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]
//输出：
//[null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]
//
//解释：
//DinnerPlates D = DinnerPlates(2);  // 初始化，栈最大容量 capacity = 2
//D.push(1);
//D.push(2);
//D.push(3);
//D.push(4);
//D.push(5);         // 栈的现状为：    2  4
//                                    1  3  5
//                                    ﹈ ﹈ ﹈
//D.popAtStack(0);   // 返回 2。栈的现状为：      4
//                                          1  3  5
//                                          ﹈ ﹈ ﹈
//D.push(20);        // 栈的现状为：  20  4
//                                   1  3  5
//                                   ﹈ ﹈ ﹈
//D.push(21);        // 栈的现状为：  20  4 21
//                                   1  3  5
//                                   ﹈ ﹈ ﹈
//D.popAtStack(0);   // 返回 20。栈的现状为：       4 21
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
//D.popAtStack(2);   // 返回 21。栈的现状为：       4
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
//D.pop()            // 返回 5。栈的现状为：        4
//                                            1  3
//                                            ﹈ ﹈
//D.pop()            // 返回 4。栈的现状为：    1  3
//                                           ﹈ ﹈
//D.pop()            // 返回 3。栈的现状为：    1
//                                           ﹈
//D.pop()            // 返回 1。现在没有栈。
//D.pop()            // 返回 -1。仍然没有栈。
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 20000
// 1 <= val <= 20000
// 0 <= index <= 100000
// 最多会对 push，pop，和 popAtStack 进行 200000 次调用。
//
//
// 👍 63 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type DinnerPlates struct {
	capacity int
	stack    []int // 模拟栈
	top      []int // 每个栈栈顶元素的位置
	popPos   []int // 有序数组保存被popAtStack删除的位置
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity: capacity, stack: []int{}, top: []int{}, popPos: []int{}}
}

func (this *DinnerPlates) Push(val int) {
	if len(this.popPos) == 0 {
		// 之前没有删除过元素，则直接添加到栈中
		pos := len(this.stack)
		this.stack = append(this.stack, val)
		if pos%this.capacity == 0 {
			this.top = append(this.top, 0)
		} else {
			idx := len(this.top) - 1
			this.top[idx] = this.top[idx] + 1
		}
	} else {
		// 之前删除过元素，则直接保存到之前删除的最小位置上
		pos := this.popPos[0]
		this.popPos = this.popPos[1:]
		this.stack[pos] = val
		idx := pos / this.capacity
		this.top[idx] = this.top[idx] + 1
	}
}

func (this *DinnerPlates) Pop() int {
	for len(this.stack) > 0 && len(this.popPos) > 0 && this.popPos[len(this.popPos)-1] == len(this.stack)-1 {
		this.stack = this.stack[:len(this.stack)-1]
		pos := this.popPos[len(this.popPos)-1]
		this.popPos = this.popPos[:len(this.popPos)-1]
		if pos%this.capacity == 0 {
			this.top = this.top[:len(this.top)-1]
		}
	}
	if len(this.stack) == 0 {
		return -1
	} else {
		pos := len(this.stack) - 1
		val := this.stack[pos]
		this.stack = this.stack[:pos]
		if pos%this.capacity == 0 && len(this.top) > 0 {
			this.top = this.top[:len(this.top)-1]
		} else if len(this.top) > 0 {
			this.top[len(this.top)-1] -= 1
		}
		return val
	}
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index >= len(this.top) {
		return -1
	}
	stackTop := this.top[index]
	if stackTop < 0 {
		return -1
	}
	this.top[index] = stackTop - 1
	pos := index*this.capacity + stackTop
	i := sort.SearchInts(this.popPos, pos)
	this.popPos = append(this.popPos, 0)
	copy(this.popPos[i+1:], this.popPos[i:])
	this.popPos[i] = pos
	return this.stack[pos]
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
