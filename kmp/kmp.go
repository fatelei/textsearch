// Package kmp implements kmp algorithm.

package kmp

// Get same sub-string of parted, return the number of same sub-string.
func findSameSubstr(parted string) int {
	length := len(parted)

	r2l := length - 1
	count := 0

	for i := 0; i < length-1; i++ {
		if parted[0:i+1] == parted[r2l-i:length] {
			tmp := i + 1
			if count < tmp {
				count = tmp
			}
		}
	}
	return count
}

// Bulld moving table, return an array.
func buildTable(keyword string) []int {
	length := len(keyword)
	table := make([]int, length)

	table[0] = 0
	for i := 1; i < length-1; i++ {
		table[i] = findSameSubstr(keyword[0 : i+1])
	}

	return table
}

// Search using kmp algorithm to find keyword in content or not.
func Search(content string, keyword string) bool {
	table := buildTable(keyword)
	contentLength := len(content)
	keywordLength := len(keyword)

	matchedLength := 0

	for i := 0; i < contentLength; {
		for j := 0; j < keywordLength; j++ {
			if content[i] != keyword[j] {
				if matchedLength > 0 {
					i = i + keywordLength - table[j-1]
					if i > contentLength {
						return false
					}
				} else {
					i++
				}
				break
			} else {
				i++
				matchedLength++
			}
		}

		if matchedLength == keywordLength {
			return true
		}

		matchedLength = 0
	}
	return false
}
