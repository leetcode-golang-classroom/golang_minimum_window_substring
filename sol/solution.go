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
