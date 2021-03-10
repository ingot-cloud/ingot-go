package pathmatcher

import (
	"strings"
)

// Ant 路径匹配
type Ant struct {
	tokenizedPatternCache map[string][]string
}

// NewAntPathMatcher 实例化
func NewAntPathMatcher() *Ant {
	return &Ant{
		tokenizedPatternCache: make(map[string][]string),
	}
}

// Match 匹配
// /** 匹配所有
// /a/** 匹配 /a/b, /a, /a/b/c
// /a/*/b 匹配 /a/c/b
// /a/**/b 匹配 /a/c/b, /a/cs/d/b
// 该方法不处理  /a/**/b/**/c多个**的情况，该路径视为/a/**/c
func (a *Ant) Match(pattern string, path string) bool {
	pattDirs := a.tokenizePattern(pattern)
	pathDirs := a.tokenizePath(path)

	pattIdxStart := 0
	pattIdxEnd := len(pattDirs) - 1
	pathIdxStart := 0
	pathIdxEnd := len(pathDirs) - 1

	for {
		if pattIdxStart > pattIdxEnd || pathIdxStart > pathIdxEnd {
			break
		}

		pattDir := pattDirs[pattIdxStart]
		if pattDir == "**" {
			break
		}

		if !a.matchStrings(pattDir, pathDirs[pathIdxStart]) {
			return false
		}

		pattIdxStart++
		pathIdxStart++
	}

	// 判断路径是否校验完毕
	// 如果路径校验完毕，pathIdxStart 大于 pathIdxEnd
	if pathIdxStart > pathIdxEnd {
		// 如果pattern校验完毕，则匹配
		if pattIdxStart > pattIdxEnd {
			return true
		}
		// 如果pattern索引到了最后一位，并且最后一位是*，那么匹配
		if pattIdxStart == pattIdxEnd && pattDirs[pattIdxStart] == "*" {
			return true
		}
		// 如果pattern索引没有到最后一位，那么从当前pattIdxStart开始依次判断是不是**，如果不是则不匹配
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if pattDirs[i] != "**" {
				return false
			}
		}
		return true
	}

	// 如果路径没有校验完毕，但是pattDirs索引结束，那么不匹配
	if pattIdxStart > pattIdxEnd {
		return false
	}

	// 如果路径没有校验完毕，pattDirs索引也没结束，
	// 那么是在pattDirs匹配到**后跳出循环的，pattern可能为 /a/** , /a/**/b
	// 由于路径也没有校验完毕，所以需要匹配pattern中**后面的索引
	// 从后往前依次校验
	for {
		if pattIdxStart > pattIdxEnd || pathIdxStart > pathIdxEnd {
			break
		}
		pattDir := pattDirs[pattIdxEnd]
		if pattDir == "**" {
			break
		}

		if !a.matchStrings(pattDir, pathDirs[pathIdxEnd]) {
			return false
		}

		pattIdxEnd--
		pathIdxEnd--
	}

	// 判断path校验完成
	if pathIdxStart > pathIdxEnd {
		// path校验完成，如果有剩余的pattDirs索引，那么必须为**才能匹配通过
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if pattDirs[i] != "**" {
				return false
			}
		}
		return true
	}

	// 如果path没有校验完
	// 判断pattDirs索引pattIdxStart是否等于pattIdxEnd
	// 由于进入到倒序校验的条件是pattIdxStart索引对应的值为**
	// 所以pattIdxEnd索引在递减的时候永远不可能小于pattIdxStart，只有在等于的时候跳出循环
	// 所以pattDirs的索引有如下几种情况
	// 1. pattIdxStart == pattIdxEnd，通过值等于**跳出循环，该情况匹配成功，
	// 示例 pattern = /a/**/b, url = /a/c/d/b, url比pattern长导致pattIdxStart和pattIdxEnd索引相同跳出循环
	if pattIdxStart == pattIdxEnd {
		return true
	}

	// 2. pattIdxStart != pattIdxEnd，由于前置条件已经跳出循环，pathIdxStart <= pathIdxEnd
	// 所以此时跳出循环的条件为 pattDirs[pattIdxEnd] == "**", 所以pattern可能为/a/**/b/c/**/d
	// 所以此时pattern中间包含多个**, pattDirs[pattIdxStart] == "**", pattDirs[pattIdxEnd] == "**"目前剩余中间部分 **/c/** 或者 **/** 需要校验
	for {
		if pattIdxStart == pattIdxEnd || pathIdxStart > pathIdxEnd {
			break
		}

		// 保存从pattIdxStart下一个位置开始，寻找**的索引
		pattIdxTemp := -1
		for i := pattIdxStart + 1; i <= pattIdxEnd; i++ {
			if pattDirs[i] == "**" {
				pattIdxTemp = i
				break
			}
		}

		if pattIdxTemp == pattIdxStart+1 {
			// 当前为连续**，也就是 **/**，直接跳过到下次循环
			pattIdxStart++
			continue
		}

		// pattIdxTemp 和 pattIdxStart不连续，那么需要校验 pattIdxTemp 和 pattIdxStart中间的值
		// 由于pattIdxTemp和pattIdxStart对应的值都是**，那么需要在pathIdxStart到pathIdxEnd中查询完全匹配pattIdxTemp 和 pattIdxStart中间的值
		// 比如 pattIdxTemp 和 pattIdxStart中间的值为 /a/b/c 那么需要在 pathIdxStart和pathIdxEnd直接查询完全匹配的值
		patLength := pattIdxTemp - pattIdxStart - 1 // 减1是因为pattIdxTemp对应的值是**，不需要校验，这里只是计算两个**直接的值的长度
		strLength := pathIdxEnd - pathIdxStart + 1  // 加一是因为校验的值包含两边的值， pathIdxStart和pathIdxEnd
		foundIdx := -1

		// 需要校验的字符串长度
		len := strLength - patLength
	strLoop:
		for i := 0; i <= len; i++ {
			for j := 0; j < patLength; j++ {
				subPat := pattDirs[pattIdxStart+j+1] // 加1是因为pattIdxStart对应的值是**，直接从下一个开始校验
				subStr := pathDirs[pathIdxStart+i+j]
				if !a.matchStrings(subPat, subStr) {
					continue strLoop
				}
			}
			foundIdx = pathIdxStart + i
			break
		}
		// foundIdx = -1 有两种情况
		// 第一种是没有找到匹配的索引
		// 第二种是 patLength大于strLength，所以校验的长度不够不匹配
		if foundIdx == -1 {
			return false
		}

		// foundIdx != -1, 代表子pattern匹配成功，继续循环校验
		// 设置 pattIdxStart=pattIdxTemp，pathIdxStart=foundIdx + patLength 并且继续循环判断
		pattIdxStart = pattIdxTemp
		pathIdxStart = foundIdx + patLength
	}

	// 执行到这一步有两种情况，一个是pattIdxStart = pattIdxEnd，还有就是pathIdxStart>pathIdxEnd(也就是字符串校验完成)
	// 增加下面这步循环校验是为了避免path校验完后，pattern还有剩余需要校验的内容，如果剩余内容都为**那么校验通过，否则不匹配
	for i := pattIdxStart; i <= pattIdxEnd; i++ {
		if pattDirs[i] != "**" {
			return false
		}
	}

	return true
}

// 简单处理，目前只匹配相等和*
func (a *Ant) matchStrings(pattern, target string) bool {
	if pattern == "*" {
		return true
	}
	return pattern == target
}

func (a *Ant) tokenizePattern(pattern string) []string {
	var tokenized []string

	tokenized = a.tokenizedPatternCache[pattern]
	if len(tokenized) != 0 {
		return tokenized
	}

	tokenized = a.tokenizePath(pattern)
	a.tokenizedPatternCache[pattern] = tokenized

	return tokenized
}

func (a *Ant) tokenizePath(path string) []string {
	var result []string

	arr := strings.Split(path, "/")
	for _, dir := range arr {
		if trim := strings.TrimSpace(dir); trim != "" {
			result = append(result, trim)
		}
	}

	return result
}
