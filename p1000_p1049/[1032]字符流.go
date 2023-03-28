package main

//设计一个算法：接收一个字符流，并检查这些字符的后缀是否是字符串数组 words 中的一个字符串。
//
// 例如，words = ["abc", "xyz"] 且字符流中逐个依次加入 4 个字符 'a'、'x'、'y' 和 'z' ，你所设计的算法应当可以检测到
// "axyz" 的后缀 "xyz" 与 words 中的字符串 "xyz" 匹配。
//
// 按下述要求实现 StreamChecker 类：
//
//
// StreamChecker(String[] words) ：构造函数，用字符串数组 words 初始化数据结构。
// boolean query(char letter)：从字符流中接收一个新字符，如果字符流中的任一非空后缀能匹配 words 中的某一字符串，返回
//true ；否则，返回 false。
//
//
//
//
// 示例：
//
//
//输入：
//["StreamChecker", "query", "query", "query", "query", "query", "query",
//"query", "query", "query", "query", "query", "query"]
//[[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"], ["h"],
//["i"], ["j"], ["k"], ["l"]]
//输出：
//[null, false, false, false, true, false, true, false, false, false, false,
//false, true]
//
//解释：
//StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
//streamChecker.query("a"); // 返回 False
//streamChecker.query("b"); // 返回 False
//streamChecker.query("c"); // 返回n False
//streamChecker.query("d"); // 返回 True ，因为 'cd' 在 words 中
//streamChecker.query("e"); // 返回 False
//streamChecker.query("f"); // 返回 True ，因为 'f' 在 words 中
//streamChecker.query("g"); // 返回 False
//streamChecker.query("h"); // 返回 False
//streamChecker.query("i"); // 返回 False
//streamChecker.query("j"); // 返回 False
//streamChecker.query("k"); // 返回 False
//streamChecker.query("l"); // 返回 True ，因为 'kl' 在 words 中
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 2000
// 1 <= words[i].length <= 200
// words[i] 由小写英文字母组成
// letter 是一个小写英文字母
// 最多调用查询 4 * 10⁴ 次
//
//
// 👍 126 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type TrieNode struct {
	Child []*TrieNode
	Fail  *TrieNode
	End   bool
}

type StreamChecker struct {
	root *TrieNode
	temp *TrieNode
}

func Constructor(words []string) StreamChecker {
	root := &TrieNode{Child: make([]*TrieNode, 26)}
	for _, word := range words {
		cur := root
		for _, ch := range word {
			if cur.Child[ch-'a'] == nil {
				cur.Child[ch-'a'] = &TrieNode{Child: make([]*TrieNode, 26)}
			}
			cur = cur.Child[ch-'a']
		}
		cur.End = true
	}
	root.Fail = root
	queue := []*TrieNode{}
	for i := 0; i < 26; i++ {
		if root.Child[i] != nil {
			root.Child[i].Fail = root
			queue = append(queue, root.Child[i])
		} else {
			root.Child[i] = root
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		cur.End = cur.End || cur.Fail.End
		for i := 0; i < 26; i++ {
			if cur.Child[i] != nil {
				cur.Child[i].Fail = cur.Fail.Child[i]
				queue = append(queue, cur.Child[i])
			} else {
				cur.Child[i] = cur.Fail.Child[i]
			}
		}
		queue = queue[1:]
	}
	return StreamChecker{root: root, temp: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.temp = this.temp.Child[letter-'a']
	return this.temp.End
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
//leetcode submit region end(Prohibit modification and deletion)
