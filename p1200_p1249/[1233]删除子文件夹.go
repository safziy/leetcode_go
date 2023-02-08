package main

import (
	"strings"
)

//你是一位系统管理员，手里有一份文件夹列表 folder，你的任务是要删除该列表中的所有 子文件夹，并以 任意顺序 返回剩下的文件夹。
//
// 如果文件夹 folder[i] 位于另一个文件夹 folder[j] 下，那么 folder[i] 就是 folder[j] 的 子文件夹 。
//
// 文件夹的「路径」是由一个或多个按以下格式串联形成的字符串：'/' 后跟一个或者多个小写英文字母。
//
//
// 例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路径，而空字符串和 "/" 不是。
//
//
//
//
// 示例 1：
//
//
//输入：folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
//输出：["/a","/c/d","/c/f"]
//解释："/a/b" 是 "/a" 的子文件夹，而 "/c/d/e" 是 "/c/d" 的子文件夹。
//
//
// 示例 2：
//
//
//输入：folder = ["/a","/a/b/c","/a/b/d"]
//输出：["/a"]
//解释：文件夹 "/a/b/c" 和 "/a/b/d" 都会被删除，因为它们都是 "/a" 的子文件夹。
//
//
// 示例 3：
//
//
//输入: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
//输出: ["/a/b/c","/a/b/ca","/a/b/d"]
//
//
//
// 提示：
//
//
// 1 <= folder.length <= 4 * 10⁴
// 2 <= folder[i].length <= 100
// folder[i] 只包含小写字母和 '/'
// folder[i] 总是以字符 '/' 起始
// 每个文件夹名都是 唯一 的
//
//
// Related Topics 字典树 数组 字符串 👍 117 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type trie struct {
	child map[string]*trie
}

func addPath(node *trie, paths []string) {
	if len(paths) == 0 {
		node.child = nil
		return
	}
	if node.child != nil {
		c, ok := node.child[paths[0]]
		if !ok {
			c = &trie{child: make(map[string]*trie)}
			node.child[paths[0]] = c
		}
		addPath(c, paths[1:])
	}
}

func removeSubfolders(folder []string) []string {
	root := &trie{child: make(map[string]*trie)}
	for _, fullPath := range folder {
		paths := strings.Split(fullPath, "/")
		addPath(root, paths[1:])
	}
	var result []string
	var dfs func(node *trie, paths []string)
	dfs = func(node *trie, paths []string) {
		if node.child == nil {
			result = append(result, strings.Join(paths, "/"))
			return
		}
		for path, child := range node.child {
			dfs(child, append(paths, path))
		}
	}
	dfs(root, []string{""})
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
