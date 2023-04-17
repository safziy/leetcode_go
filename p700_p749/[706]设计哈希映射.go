package main

import "container/list"

//不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
//
// 实现 MyHashMap 类：
//
//
// MyHashMap() 用空映射初始化对象
// void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，
//则更新其对应的值 value 。
// int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
// void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
//
//
//
//
// 示例：
//
//
//输入：
//["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
//[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
//输出：
//[null, null, null, 1, -1, null, 1, null, -1]
//
//解释：
//MyHashMap myHashMap = new MyHashMap();
//myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
//myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
//myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
//myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
//myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
//
//
//
//
// 提示：
//
//
// 0 <= key, value <= 10⁶
// 最多调用 10⁴ 次 put、get 和 remove 方法
//
//
// 👍 360 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
const base = 769

type entry struct{ key, value int }

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (this *MyHashMap) hash(key int) int { return key % base }

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[h].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
//leetcode submit region end(Prohibit modification and deletion)
