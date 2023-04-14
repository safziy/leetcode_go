# 前缀函数与 KMP 算法

## 字符串前缀和后缀定义

关于字符串前缀、真前缀，后缀、真后缀的定义详见 [字符串基础](字符串基础.md)

## 前缀函数

### 定义

给定一个长度为n的字符串s，其 **前缀函数** 被定义为一个长度为n的数组π。 其中π[i]的定义是：

1.  如果子串s[0...i] 有一对相等的真前缀与真后缀：s[0...k-1]和s[i - (k - 1)...i]，那么π[i]就是这个相等的真前缀（或者真后缀，因为它们相等：）的长度，也就是π[i]=k
2. 如果不止有一对相等的，那么π[i]就是其中最长的那一对的长度；
3.  如果没有相等的，那么π[i]=0。

简单来说π[i]就是，子串s[0...i]最长的相等的真前缀与真后缀的长度。

用数学语言描述如下："\pi[i] = \max_{k = 0 \dots i}\{k: s[0 \dots k - 1] = s[i - &#40;k - 1&#41; \dots i]\}"

π[i]=$\max \limits_{k 0 \to i} \lgroup k:s[0...k-1] = s[i-(k-1)...i] \rgroup$

特别地，规定π[i]=0。

### 过程[](#过程 "Permanent link")

举例来说，对于字符串 `abcabcd`，

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[0]=0")，因为 `a` 没有真前缀和真后缀，根据规定为 0

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[1]=0")，因为 `ab` 无相等的真前缀和真后缀

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[2]=0")，因为 `abc` 无相等的真前缀和真后缀

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[3]=1")，因为 `abca` 只有一对相等的真前缀和真后缀：`a`，长度为 1

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[4]=2")，因为 `abcab` 相等的真前缀和真后缀只有 `ab`，长度为 2

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[5]=3")，因为 `abcabc` 相等的真前缀和真后缀只有 `abc`，长度为 3

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[6]=0")，因为 `abcabcd` 无相等的真前缀和真后缀

同理可以计算字符串 `aabaaab` 的前缀函数为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "[0, 1, 0, 1, 2, 2, 3]")。

## 计算前缀函数的朴素算法[](#计算前缀函数的朴素算法 "Permanent link")

### 过程[](#过程_1 "Permanent link")

一个直接按照定义计算前缀函数的算法流程：

- 在一个循环中以

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i = 1\to n - 1") 的顺序计算前缀函数

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 的值（

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[0]") 被赋值为

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "0")）。
- 为了计算当前的前缀函数值

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]")，我们令变量

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 从最大的真前缀长度

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 开始尝试。
- 如果当前长度下真前缀和真后缀相等，则此时长度为

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]")，否则令 j 自减 1，继续匹配，直到

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j=0")。
- 如果

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j = 0") 并且仍没有任何一次匹配，则置

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i] = 0") 并移至下一个下标

  ![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i + 1")。

实现

具体实现如下：

C++Python

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>``` | ```<br><br><br>`vector<int> <br><br> prefix_function(string <br><br> s) <br><br> {<br><br> <br><br> int <br><br> n <br><br> = <br><br> (int)s.length();<br><br> <br><br> vector<int> <br><br> pi(n);<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> 1; <br><br> i <br><br> < <br><br> n; <br><br> i++)<br><br> <br><br> for <br><br> (int <br><br> j <br><br> = <br><br> i; <br><br> j <br><br> >= <br><br> 0; <br><br> j--)<br><br> <br><br> if <br><br> (s.substr(0, <br><br> j) <br><br> == <br><br> s.substr(i <br><br> - <br><br> j <br><br> + <br><br> 1, <br><br> j)) <br><br> {<br><br> <br><br> pi[i] <br><br> = <br><br> j;<br><br> <br><br> break;<br><br> <br><br> }<br><br> <br><br> return <br><br> pi;<br><br>}` <br>``` |

注：

- `string substr (size_t pos = 0, size_t len = npos) const;`

显见该算法的时间复杂度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n^3&#41;")，具有很大的改进空间。

## 计算前缀函数的高效算法[](#计算前缀函数的高效算法 "Permanent link")

### 第一个优化[](#第一个优化 "Permanent link")

第一个重要的观察是 **相邻的前缀函数值至多增加

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "1")** 。

参照下图所示，只需如此考虑：当取一个尽可能大的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i+1]") 时，必然要求新增的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i+1]") 也与之对应的字符匹配，即

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i+1]=s[\pi[i]]"), 此时

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i+1] = \pi[i]+1")。

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\underbrace{\overbrace{s_0 ~ s_1 ~ s_2}^{\pi[i] = 3} ~ s_3}_{\pi[i+1] = 4} ~ \dots ~ \underbrace{\overbrace{s_{i-2} ~ s_{i-1} ~ s_{i}}^{\pi[i] = 3} ~ s_{i+1}}_{\pi[i+1] = 4}")

所以当移动到下一个位置时，前缀函数的值要么增加一，要么维持不变，要么减少。

实现

此时的改进的算法为：

C++Python

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>``` | ```<br><br><br>`vector<int> <br><br> prefix_function(string <br><br> s) <br><br> {<br><br> <br><br> int <br><br> n <br><br> = <br><br> (int)s.length();<br><br> <br><br> vector<int> <br><br> pi(n);<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> 1; <br><br> i <br><br> < <br><br> n; <br><br> i++)<br><br> <br><br> for <br><br> (int <br><br> j <br><br> = <br><br> pi[i <br><br> - <br><br> 1] <br><br> + <br><br> 1; <br><br> j <br><br> >= <br><br> 0; <br><br> j--) <br><br> // improved: j=i => j=pi[i-1]+1<br> <br><br> if <br><br> (s.substr(0, <br><br> j) <br><br> == <br><br> s.substr(i <br><br> - <br><br> j <br><br> + <br><br> 1, <br><br> j)) <br><br> {<br><br> <br><br> pi[i] <br><br> = <br><br> j;<br><br> <br><br> break;<br><br> <br><br> }<br><br> <br><br> return <br><br> pi;<br><br>}` <br>``` | 

在这个初步改进的算法中，在计算每个

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 时，最好的情况是第一次字符串比较就完成了匹配，也就是说基础的字符串比较次数是 `n-1` 次。

而由于存在 `j = pi[i-1]+1`（`pi[0]=0`）对于最大字符串比较次数的限制，可以看出每次只有在最好情况才会为字符串比较次数的上限积累 1，而每次超过一次的字符串比较消耗的是之后次数的增长空间。

由此我们可以得出字符串比较次数最多的一种情况：至少 `1` 次字符串比较次数的消耗和最多 `n-2` 次比较次数的积累，此时字符串比较次数为 `n-1 + n-2 = 2n-3`。

可见经过此次优化，计算前缀函数只需要进行

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n&#41;") 次字符串比较，总复杂度降为了

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n^2&#41;")。

### 第二个优化[](#第二个优化 "Permanent link")

在第一个优化中，我们讨论了计算

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i+1]") 时的最好情况：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i+1]=s[\pi[i]]")，此时

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i+1] = \pi[i]+1")。现在让我们沿着这个思路走得更远一点：讨论当

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i+1] \neq s[\pi[i]]") 时如何跳转。

失配时，我们希望找到对于子串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0\dots i]")，仅次于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 的第二长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j")，使得在位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的前缀性质仍得以保持，也即

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0 \dots j - 1] = s[i - j + 1 \dots i]")：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\overbrace{\underbrace{s_0 ~ s_1}_j ~ s_2 ~ s_3}^{\pi[i]} ~ \dots ~ \overbrace{s_{i-3} ~ s_{i-2} ~ \underbrace{s_{i-1} ~ s_{i}}_j}^{\pi[i]} ~ s_{i+1}")

如果我们找到了这样的长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j")，那么仅需要再次比较

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i + 1]") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[j]")。如果它们相等，那么就有

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i + 1] = j + 1")。否则，我们需要找到子串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0\dots i]") 仅次于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 的第二长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j^{&#40;2&#41;}")，使得前缀性质得以保持，如此反复，直到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j = 0")。如果

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i + 1] \neq s[0]")，则

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i + 1] = 0")。

观察上图可以发现，因为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0\dots \pi[i]-1] = s[i-\pi[i]+1\dots i]")，所以对于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0\dots i]") 的第二长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j")，有这样的性质：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0 \dots j - 1] = s[i - j + 1 \dots i]= s[\pi[i]-j\dots \pi[i]-1]")

也就是说

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 等价于子串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[\pi[i]-1]") 的前缀函数值，即

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j=\pi[\pi[i]-1]")。同理，次于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 的第二长度等价于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[j-1]") 的前缀函数值，

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j^{&#40;2&#41;}=\pi[j-1]")

显然我们可以得到一个关于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 的状态转移方程：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j^{&#40;n&#41;}=\pi[j^{&#40;n-1&#41;}-1], \ \ &#40;j^{&#40;n-1&#41;}>0&#41;")

### 最终算法[](#最终算法 "Permanent link")

所以最终我们可以构建一个不需要进行任何字符串比较，并且只进行

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n&#41;") 次操作的算法。

而且该算法的实现出人意料的短且直观：

实现

C++Python

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>``` | ```<br><br><br>`vector<int> <br><br> prefix_function(string <br><br> s) <br><br> {<br><br> <br><br> int <br><br> n <br><br> = <br><br> (int)s.length();<br><br> <br><br> vector<int> <br><br> pi(n);<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> 1; <br><br> i <br><br> < <br><br> n; <br><br> i++) <br><br> {<br><br> <br><br> int <br><br> j <br><br> = <br><br> pi[i <br><br> - <br><br> 1];<br><br> <br><br> while <br><br> (j <br><br> > <br><br> 0 <br><br> && <br><br> s[i] <br><br> != <br><br> s[j]) <br><br> j <br><br> = <br><br> pi[j <br><br> - <br><br> 1];<br><br> <br><br> if <br><br> (s[i] <br><br> == <br><br> s[j]) <br><br> j++;<br><br> <br><br> pi[i] <br><br> = <br><br> j;<br><br> <br><br> }<br><br> <br><br> return <br><br> pi;<br><br>}` <br>``` | 

这是一个 **在线** 算法，即其当数据到达时处理它——举例来说，你可以一个字符一个字符的读取字符串，立即处理它们以计算出每个字符的前缀函数值。该算法仍然需要存储字符串本身以及先前计算过的前缀函数值，但如果我们已经预先知道该字符串前缀函数的最大可能取值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "M")，那么我们仅需要存储该字符串的前

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "M + 1") 个字符以及对应的前缀函数值。

## 应用[](#应用 "Permanent link")

### 在字符串中查找子串：Knuth-Morris-Pratt 算法[](#在字符串中查找子串knuth-morris-pratt-算法 "Permanent link")

该算法由 Knuth、Pratt 和 Morris 在 1977 年共同发布<sup>[\[1\]](https://epubs.siam.org/doi/abs/10.1137/0206024)</sup>。

该任务是前缀函数的一个典型应用。

#### 过程[](#过程_2 "Permanent link")

给定一个文本

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 和一个字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")，我们尝试找到并展示

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 中的所有出现（occurrence）。

为了简便起见，我们用

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 表示字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的长度，用

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "m") 表示文本

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的长度。

我们构造一个字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t")，其中

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\#") 为一个既不出现在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 中也不出现在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 中的分隔符。接下来计算该字符串的前缀函数。现在考虑该前缀函数除去最开始

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n + 1") 个值（即属于字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 和分隔符的函数值）后其余函数值的意义。根据定义，

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 为右端点在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 且同时为一个前缀的最长真子串的长度，具体到我们的这种情况下，其值为与

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的前缀相同且右端点位于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的最长子串的长度。由于分隔符的存在，该长度不可能超过

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n")。而如果等式

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i] = n") 成立，则意味着

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 完整出现在该位置（即其右端点位于位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i")）。注意该位置的下标是对字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t") 而言的。

因此如果在某一位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 有

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i] = n") 成立，则字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i - &#40;n - 1&#41; - &#40;n + 1&#41; = i - 2n") 处出现。

正如在前缀函数的计算中已经提到的那样，如果我们知道前缀函数的值永远不超过一特定值，那么我们不需要存储整个字符串以及整个前缀函数，而只需要二者开头的一部分。在我们这种情况下这意味着只需要存储字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \#") 以及相应的前缀函数值即可。我们可以一次读入字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的一个字符并计算当前位置的前缀函数值。

因此 Knuth-Morris-Pratt 算法（简称 KMP 算法）用

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n + m&#41;") 的时间以及

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n&#41;") 的内存解决了该问题。

实现

C++Python

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>``` | ```<br><br><br>`vector<int> <br><br> find_occurrences(string <br><br> text, <br><br> string <br><br> pattern) <br><br> {<br><br> <br><br> string <br><br> cur <br><br> = <br><br> pattern <br><br> + <br><br> '#' <br><br> + <br><br> text;<br><br> <br><br> int <br><br> sz1 <br><br> = <br><br> text.size(), <br><br> sz2 <br><br> = <br><br> pattern.size();<br><br> <br><br> vector<int> <br><br> v;<br><br> <br><br> vector<int> <br><br> lps <br><br> = <br><br> prefix_function(cur);<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> sz2 <br><br> + <br><br> 1; <br><br> i <br><br> <= <br><br> sz1 <br><br> + <br><br> sz2; <br><br> i++) <br><br> {<br><br> <br><br> if <br><br> (lps[i] <br><br> == <br><br> sz2)<br><br> <br><br> v.push_back(i <br><br> - <br><br> 2 <br><br> * <br><br> sz2);<br><br> <br><br> }<br><br> <br><br> return <br><br> v;<br><br>}` <br>``` | 

### 字符串的周期[](#字符串的周期 "Permanent link")

对字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "0 < p \le |s|")，若

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[i] = s[i+p]") 对所有

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i \in [0, |s| - p - 1]") 成立，则称

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的周期。

对字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "0 \le r < |s|")，若

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的前缀和长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的后缀相等，就称

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的前缀是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的 border。

由

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 有长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的 border 可以推导出

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "|s|-r") 是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的周期。

根据前缀函数的定义，可以得到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 所有的 border 长度，即

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n-1],\pi[\pi[n-1]-1], \ldots")。<sup>[2](#fn:ref1)</sup>

所以根据前缀函数可以在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n&#41;") 的时间内计算出

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 所有的周期。其中，由于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n-1]") 是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 最长 border 的长度，所以

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n - \pi[n-1]") 是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的最小周期。

### 统计每个前缀的出现次数[](#统计每个前缀的出现次数 "Permanent link")

在该节我们将同时讨论两个问题。给定一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 的字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")，在问题的第一个变种中我们希望统计每个前缀

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0 \dots i]") 在同一个字符串的出现次数，在问题的第二个变种中我们希望统计每个前缀

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s[0 \dots i]") 在另一个给定字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 中的出现次数。

首先让我们来解决第一个问题。考虑位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的前缀函数值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]")。根据定义，其意味着字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 的前缀在位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 出现并以

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 为右端点，同时不存在一个更长的前缀满足前述定义。与此同时，更短的前缀可能以该位置为右端点。容易看出，我们遇到了在计算前缀函数时已经回答过的问题：给定一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 的前缀，同时其也是一个右端点位于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的后缀，下一个更小的前缀长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k < j") 是多少？该长度的前缀需同时也是一个右端点为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的后缀。因此以位置

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 为右端点，有长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]") 的前缀，有长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[\pi[i] - 1]") 的前缀，有长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[\pi[\pi[i] - 1] - 1]") 的前缀，等等，直到长度变为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "0")。故而我们可以通过下述方式计算答案。

实现

C++Python

|     |     |
| --- | --- |
| ```<br><br><br>1<br>2<br>3<br>4<br>``` | ```<br><br><br>`vector<int> <br><br> ans(n <br><br> + <br><br> 1);<br><br>for <br><br> (int <br><br> i <br><br> = <br><br> 0; <br><br> i <br><br> < <br><br> n; <br><br> i++) <br><br> ans[pi[i]]++;<br><br>for <br><br> (int <br><br> i <br><br> = <br><br> n <br><br> - <br><br> 1; <br><br> i <br><br> > <br><br> 0; <br><br> i--) <br><br> ans[pi[i <br><br> - <br><br> 1]] <br><br> += <br><br> ans[i];<br><br>for <br><br> (int <br><br> i <br><br> = <br><br> 0; <br><br> i <br><br> <= <br><br> n; <br><br> i++) <br><br> ans[i]++;` <br>``` | 

#### 解释[](#解释 "Permanent link")

在上述代码中我们首先统计每个前缀函数值在数组

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi") 中出现了多少次，然后再计算最后答案：如果我们知道长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 的前缀出现了恰好

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\text{ans}[i]") 次，那么该值必须被叠加至其最长的既是后缀也是前缀的子串的出现次数中。在最后，为了统计原始的前缀，我们对每个结果加

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "1")。

现在考虑第二个问题。我们应用来自 Knuth-Morris-Pratt 的技巧：构造一个字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t") 并计算其前缀函数。与第一个问题唯一的不同之处在于，我们只关心与字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 相关的前缀函数值，即

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i \ge n + 1") 的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[i]")。有了这些值之后，我们可以同样应用在第一个问题中的算法来解决该问题。

### 一个字符串中本质不同子串的数目[](#一个字符串中本质不同子串的数目 "Permanent link")

给定一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 的字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")，我们希望计算其本质不同子串的数目。

我们将迭代的解决该问题。换句话说，在知道了当前的本质不同子串的数目的情况下，我们要找出一种在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 末尾添加一个字符后重新计算该数目的方法。

令

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 为当前

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的本质不同子串数量。我们添加一个新的字符

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "c") 至

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")。显然，会有一些新的子串以字符

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "c") 结尾。我们希望对这些以该字符结尾且我们之前未曾遇到的子串计数。

构造字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t = s + c") 并将其反转得到字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t^{\sim}")。现在我们的任务变为计算有多少

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t^{\sim}") 的前缀未在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t^{\sim}") 的其余任何地方出现。如果我们计算了

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t^{\sim}") 的前缀函数最大值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi_{\max}")，那么最长的出现在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 中的前缀其长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi_{\max}")。自然的，所有更短的前缀也出现了。

因此，当添加了一个新字符后新出现的子串数目为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "|s| + 1 - \pi_{\max}")。

所以对于每个添加的字符，我们可以在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n&#41;") 的时间内计算新子串的数目，故最终复杂度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;n^2&#41;")。

值得注意的是，我们也可以重新计算在头部添加一个字符，或者从尾或者头移除一个字符时的本质不同子串数目。

### 字符串压缩[](#字符串压缩 "Permanent link")

给定一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 的字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")，我们希望找到其最短的「压缩」表示，也即我们希望寻找一个最短的字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t")，使得

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 可以被

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的一份或多份拷贝的拼接表示。

显然，我们只需要找到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的长度即可。知道了该长度，该问题的答案即为长度为该值的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的前缀。

让我们计算

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 的前缀函数。通过使用该函数的最后一个值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n - 1]")，我们定义值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k = n - \pi[n - 1]")。我们将证明，如果

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 整除

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n")，那么

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 就是答案，否则不存在一个有效的压缩，故答案为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n")。

假定

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 可被

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 整除。那么字符串可被划分为长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 的若干块。根据前缀函数的定义，该字符串长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n - k") 的前缀等于其后缀。但是这意味着最后一个块同倒数第二个块相等，并且倒数第二个块同倒数第三个块相等，等等。作为其结果，所有块都是相等的，因此我们可以将字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 压缩至长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k")。

证明

诚然，我们仍需证明该值为最优解。实际上，如果有一个比

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 更小的压缩表示，那么前缀函数的最后一个值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n - 1]") 必定比

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n - k") 要大。因此

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 就是答案。

现在假设

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 不可以被

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 整除，我们将通过反证法证明这意味着答案为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n")<sup>[1](#fn:1)</sup>。假设其最小压缩表示

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p")（

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 整除

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n")），字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 被划分为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n / p \ge 2") 块。那么前缀函数的最后一个值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n - 1]") 必定大于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n - p")（如果等于则

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n") 可被

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 整除），也即其所表示的后缀将部分的覆盖第一个块。现在考虑字符串的第二个块。该块有两种解释：第一种为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r_0 r_1 \dots r_{p - 1}")，另一种为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r_{p - k} r_{p - k + 1} \dots r_{p - 1} r_0 r_1 \dots r_{p - k - 1}")。由于两种解释对应同一个字符串，因此可得到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 个方程组成的方程组，该方程组可简写为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r_{&#40;i + k&#41; \bmod p} = r_{i \bmod p}")，其中

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\cdot \bmod p") 表示模

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 意义下的最小非负剩余。

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\begin{gathered}
\overbrace{r_0 ~ r_1 ~ r_2 ~ r_3 ~ r_4 ~ r_5}^p ~ \overbrace{r_0 ~ r_1 ~ r_2 ~ r_3 ~ r_4 r_5}^p \\
r_0 ~ r_1 ~ r_2 ~ r_3 ~ \underbrace{\overbrace{r_0 ~ r_1 ~ r_2 ~ r_3 ~ r_4 ~ r_5}^p ~ r_0 ~ r_1}_{\pi[11] = 8}
\end{gathered}")

根据扩展欧几里得算法我们可以得到一组

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "x") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "y") 使得

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "xk + yp = \gcd&#40;k, p&#41;")。通过与等式

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "pk - kp = 0") 适当叠加我们可以得到一组

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "x' > 0") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "y' < 0") 使得

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "x'k + y'p = \gcd&#40;k, p&#41;")。这意味着通过不断应用前述方程组中的方程我们可以得到新的方程组

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r_{&#40;i + \gcd&#40;k, p&#41;&#41; \bmod p} = r_{i \bmod p}")。

由于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\gcd&#40;k, p&#41;") 整除

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p")，这意味着

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\gcd&#40;k, p&#41;") 是

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的一个周期。又因为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[n - 1] > n - p")，故有

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "n - \pi[n - 1] = k < p")，所以

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\gcd&#40;k, p&#41;") 是一个比

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 更小的

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "r") 的周期。因此字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 有一个长度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\gcd&#40;k, p&#41; < p") 的压缩表示，同

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "p") 的最小性矛盾。

综上所述，不存在一个长度小于

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 的压缩表示，因此答案为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k")。

### 根据前缀函数构建一个自动机[](#根据前缀函数构建一个自动机 "Permanent link")

让我们重新回到通过一个分隔符将两个字符串拼接的新字符串。对于字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 和

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 我们计算

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t") 的前缀函数。显然，因为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\#") 是一个分隔符，前缀函数值永远不会超过

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "|s|")。因此我们只需要存储字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \#") 和其对应的前缀函数值，之后就可以动态计算对于之后所有字符的前缀函数值：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\underbrace{s_0 ~ s_1 ~ \dots ~ s_{n-1} ~ \#}_{\text{need to store}} ~ \underbrace{t_0 ~ t_1 ~ \dots ~ t_{m-1}}_{\text{do not need to store}}")

实际上在这种情况下，知道

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的下一个字符

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "c") 以及之前位置的前缀函数值便足以计算下一个位置的前缀函数值，而不需要用到任何其它

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 的字符和对应的前缀函数值。

换句话说，我们可以构造一个 **自动机**（一个有限状态机）：其状态为当前的前缀函数值，而从一个状态到另一个状态的转移则由下一个字符确定。

因此，即使没有字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t")，我们同样可以应用构造转移表的算法构造一个转移表

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "&#40; \text { old } \pi , c &#41; \rightarrow \text { new } _ { - } \pi")：

实现

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>``` | ```<br><br><br>`void <br><br> compute_automaton(string <br><br> s, <br><br> vector<vector<int>>& <br><br> aut) <br><br> {<br><br> <br><br> s <br><br> += <br><br> '#';<br><br> <br><br> int <br><br> n <br><br> = <br><br> s.size();<br><br> <br><br> vector<int> <br><br> pi <br><br> = <br><br> prefix_function(s);<br><br> <br><br> aut.assign(n, <br><br> vector<int>(26));<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> 0; <br><br> i <br><br> < <br><br> n; <br><br> i++) <br><br> {<br><br> <br><br> for <br><br> (int <br><br> c <br><br> = <br><br> 0; <br><br> c <br><br> < <br><br> 26; <br><br> c++) <br><br> {<br><br> <br><br> int <br><br> j <br><br> = <br><br> i;<br><br> <br><br> while <br><br> (j <br><br> > <br><br> 0 <br><br> && <br><br> 'a' <br><br> + <br><br> c <br><br> != <br><br> s[j]) <br><br> j <br><br> = <br><br> pi[j <br><br> - <br><br> 1];<br><br> <br><br> if <br><br> ('a' <br><br> + <br><br> c <br><br> == <br><br> s[j]) <br><br> j++;<br><br> <br><br> aut[i][c] <br><br> = <br><br> j;<br><br> <br><br> }<br><br> <br><br> }<br><br>}` <br>``` | 

然而在这种形式下，对于小写字母表，算法的时间复杂度为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;|\Sigma|n^2&#41;")。注意到我们可以应用动态规划来利用表中已计算过的部分。只要我们从值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 变化到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\pi[j - 1]")，那么我们实际上在说转移

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "&#40;j, c&#41;") 所到达的状态同转移

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "&#40;\pi[j - 1], c&#41;") 一样，但该答案我们之前已经精确计算过了。

实现

|     |     |
| --- | --- |
| ```<br><br><br> 1<br> 2<br> 3<br> 4<br> 5<br> 6<br> 7<br> 8<br> 9<br>10<br>11<br>12<br>13<br>14<br>``` | ```<br><br><br>`void <br><br> compute_automaton(string <br><br> s, <br><br> vector<vector<int>>& <br><br> aut) <br><br> {<br><br> <br><br> s <br><br> += <br><br> '#';<br><br> <br><br> int <br><br> n <br><br> = <br><br> s.size();<br><br> <br><br> vector<int> <br><br> pi <br><br> = <br><br> prefix_function(s);<br><br> <br><br> aut.assign(n, <br><br> vector<int>(26));<br><br> <br><br> for <br><br> (int <br><br> i <br><br> = <br><br> 0; <br><br> i <br><br> < <br><br> n; <br><br> i++) <br><br> {<br><br> <br><br> for <br><br> (int <br><br> c <br><br> = <br><br> 0; <br><br> c <br><br> < <br><br> 26; <br><br> c++) <br><br> {<br><br> <br><br> if <br><br> (i <br><br> > <br><br> 0 <br><br> && <br><br> 'a' <br><br> + <br><br> c <br><br> != <br><br> s[i])<br><br> <br><br> aut[i][c] <br><br> = <br><br> aut[pi[i <br><br> - <br><br> 1]][c];<br><br> <br><br> else<br><br> <br><br> aut[i][c] <br><br> = <br><br> i <br><br> + <br><br> ('a' <br><br> + <br><br> c <br><br> == <br><br> s[i]);<br><br> <br><br> }<br><br> <br><br> }<br><br>}` <br>``` | 

最终我们可在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "O&#40;|\Sigma|n&#41;") 的时间复杂度内构造该自动机。

该自动机在什么时候有用呢？首先，记得大部分时候我们为了一个目的使用字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t") 的前缀函数：寻找字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 中的所有出现。

因此使用该自动机的最直接的好处是 **加速计算字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \# + t") 的前缀函数**。

通过构建

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s + \#") 的自动机，我们不再需要存储字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 以及其对应的前缀函数值。所有转移已经在表中计算过了。

但除此以外，还有第二个不那么直接的应用。我们可以在字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 是 **某些通过一些规则构造的巨型字符串** 时，使用该自动机加速计算。Gray 字符串，或者一个由一些短的输入串的递归组合所构造的字符串都是这种例子。

出于完整性考虑，我们来解决这样一个问题：给定一个数

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k \le 10^5")，以及一个长度

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\le 10^5") 的字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s")，我们需要计算

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在第

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 个 Gray 字符串中的出现次数。回想起 Gray 字符串以下述方式定义：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\begin{aligned}
g_1 &= \mathtt{a}\\
g_2 &= \mathtt{aba}\\
g_3 &= \mathtt{abacaba}\\
g_4 &= \mathtt{abacabadabacaba}
\end{aligned}")

由于其天文数字般的长度，在这种情况下即使构造字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t") 都是不可能的：第

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "k") 个 Gray 字符串有

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "2^k - 1") 个字符。然而我们可以在仅仅知道开头若干前缀函数值的情况下，有效计算该字符串末尾的前缀函数值。

除了自动机之外，我们同时需要计算值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "G[i][j]")：在从状态

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 开始处理

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_i") 后的自动机的状态，以及值

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[i][j]")：当从状态

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "j") 开始处理

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_i") 后，

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_i") 中的出现次数。实际上

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[i][j]") 为在执行操作时前缀函数取值为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "|s|") 的次数。易得问题的答案为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[k][0]")。

我们该如何计算这些值呢？首先根据定义，初始条件为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "G[0][j] = j") 以及

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[0][j] = 0")。之后所有值可以通过先前的值以及使用自动机计算得到。为了对某个

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 计算相应值，回想起字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_i") 由

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_{i - 1}")，字母表中第

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "i") 个字符，以及

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "g_{i - 1}") 三者拼接而成。因此自动机会途径下列状态：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\begin{gathered}
\text{mid} = \text{aut}[G[i - 1][j]][i] \\
G[i][j] = G[i - 1][\text{mid}]
\end{gathered}")

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[i][j]") 的值同样可被简单计算。

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "K[i][j] = K[i - 1][j] + [\text{mid} == |s|] + K[i - 1][\text{mid}]")

其中

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "[\cdot]") 当其中表达式取值为真时值为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "1")，否则为

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "0")。综上，我们已经可以解决关于 Gray 字符串的问题，以及一大类与之类似的问题。举例来说，应用同样的方法可以解决下列问题：给定一个字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 以及一些模式

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t_i")，其中每个模式以下列方式给出：该模式由普通字符组成，当中可能以

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t_{k}^{\text{cnt}}") 的形式递归插入先前的字符串，也即在该位置我们必须插入字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "t_k")

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\text{cnt}") 次。以下是这些模式的一个例子：

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "\begin{aligned}
t_1 &= \mathtt{abdeca} \\
t_2 &= \mathtt{abc} + t_1^{30} + \mathtt{abd} \\
t_3 &= t_2^{50} + t_1^{100} \\
t_4 &= t_2^{10} + t_3^{100}
\end{aligned}")

递归代入会使字符串长度爆炸式增长，他们的长度甚至可以达到

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "100^{100}") 的数量级。而我们必须找到字符串

![](data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7 "s") 在每个字符串中的出现次数。

该问题同样可通过构造前缀函数的自动机解决。同之前一样，我们利用先前计算过的结果对每个模式计算其转移然后相应统计答案即可。

## 练习题目[](#练习题目 "Permanent link")

- [UVA 455 "Periodic Strings"](http://uva.onlinejudge.org/index.php?option=onlinejudge&page=show_problem&problem=396)
- [UVA 11022 "String Factoring"](http://uva.onlinejudge.org/index.php?option=onlinejudge&page=show_problem&problem=1963)
- [UVA 11452 "Dancing the Cheeky-Cheeky"](http://uva.onlinejudge.org/index.php?option=onlinejudge&page=show_problem&problem=2447)
- [UVA 12604 - Caesar Cipher](https://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=4282)
- [UVA 12467 - Secret Word](https://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=3911)
- [UVA 11019 - Matrix Matcher](https://uva.onlinejudge.org/index.php?option=onlinejudge&page=show_problem&problem=1960)
- [SPOJ - Pattern Find](http://www.spoj.com/problems/NAJPF/)
- [Codeforces - Anthem of Berland](http://codeforces.com/contest/808/problem/G)
- [Codeforces - MUH and Cube Walls](http://codeforces.com/problemset/problem/471/D)