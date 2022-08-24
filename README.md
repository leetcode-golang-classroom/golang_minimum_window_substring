# golang_minimum_window_substring

Given two strings `s` and `t` of lengths `m` and `n` respectively, return *the **minimum window substring** of* `s` *such that every character in* `t` *(**including duplicates**) is included in the window. If there is no such substring, return the empty string* `""`*.*

The testcases will be generated such that the answer is **unique**.

A **substring** is a contiguous sequence of characters within the string.

## Examples

**Example 1:**

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

```

**Example 2:**

```
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

```

**Example 3:**

```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

```

**Constraints:**

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= $10^5$`
- `s` and `t` consist of uppercase and lowercase English letters.

**Follow up:**

Could you find an algorithm that runs in O(m + n) time?

## 解析

給定兩個字串 s, t

要求寫一個演算法找出 s 的子字串中包含 t 字串所有字元的最短子字串

s 的子字串 p 代表 p 中所有字元都在 s 出現過且每個字元必須跟 s 的相鄰狀態相同

所以題目要找子字串 p 必須要

 1.   包含所有 t 的字元

1. 必須是字元相連性跟原本 s 相同

 3.  必須是最短的

第一個條件可以透過 hashmap 的方式把 t 的所有字元出現次數紀錄下來比對

第二個相鄰狀態可以透過 slide-window 方式來保證

第3個條件 可以知道最理想的狀態是 剛好子字串長度跟 t 字串長度一樣

                  所以從子字串長度 == t 字串長度 且累計配對字元數相等時

                  就可以不斷更新 最小子字串長度直到把所有 s 的可能右界都走完 就可以找到最小的那個長度了

![](https://i.imgur.com/azeZeW7.png)


具體作法是大致分成三大步驟

1 在 slide-window 還沒有包含足夠多的字元時， 把右界往右

2 當符合條件具有足夠多的字元時，更新當 slide-window 比上一次符合條件小時，更新 slide-window 的 size 並且紀錄當下左右界

3 開始把左界左移，並且檢查當下 slide-window 是否有更小

具體作法如下圖

 
![](https://i.imgur.com/OfNUH08.jpg)

## 程式碼
```go
package sol

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	sFreq, tFreq := make([]int, 256), make([]int, 256)
	result := ""
	count := 0
	left, right, subLeft, subRight := 0, -1, -1, -1
	minW := sLen + 1
	if sLen < tLen || sLen == 0 || tLen == 0 {
		return result
	}
	for pos := 0; pos < tLen; pos++ {
		tFreq[t[pos]]++
	}
	for left < sLen {
		// check if right not reach end and count
		if right+1 < sLen && count < tLen {
			if sFreq[s[right+1]] < tFreq[s[right+1]] {
				count++
			}
			sFreq[s[right+1]]++
			right++
		} else {
			if right-left+1 < minW && count == tLen { // smaller window exist
				minW = right - left + 1
				subLeft = left
				subRight = right
			}
			// shift left point to right
			if sFreq[s[left]] == tFreq[s[left]] {
				count--
			}
			sFreq[s[left]]--
			left++
		}
	}
	if subLeft != -1 {
		result = s[subLeft : subRight+1]
	}
	return result
}
```
## 困難點

1. 需要理解透過 HashTable 確認字元出現個數的方法
2. 需要知道 slide-window 找尋可能子字串的方式

## Solve Point

- [x]  初始化  s_freq , t_freq 兩個 HashTable 用來各自紀錄 s, t 的字元次數出現狀況, count= 0 用來紀錄 s match 到 t 的字元個數， left = 0, right = -1,  sub_left = -1, sub_right = -1, minW = len(s)+1
- [x]  先把 t 的字元次數紀錄在 t_freq
- [x]  當 left < len(s) 做以下運算
- [x]  當 right + 1 < len(s) && count < len(t) 時 , 代表 還沒蒐集夠子字串字元 做以下運算
- [x]  更新 s_freq[s[right+1]] += 1
- [x]  當 s_freq[s[right+1]] ≤ t_freq[s[right+1]] 時 代表 該字元需要被納入 , 更新 count = count+1
- [x]  更新 right = right + 1
- [x]  當 right + 1 == len(s) || count == len(t) 時，做以下運算
- [x]  當 right - left + 1 < minW 且 count == len(t) 時 代表已經蒐集足夠字元當子字串， 更新 minW = right - left  + 1 , 更新 sub_left = left, sub_right = right
- [x]  當 s_freq[s[left]] == t_freq[s[left]] 時 , 代表要左移的字元有被算入 count 之中 , 更新 count = count - 1 來做還原
- [x]  更新 s_freq[s[left]] -= 1, left = left + 1
- [x]  當跑到最後 sub_left ≠ -1 代表有找到子字串所以回傳 s[sub_left: sub_right+1], 否則回傳 “”