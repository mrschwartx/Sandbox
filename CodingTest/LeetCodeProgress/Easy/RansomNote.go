func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, ch := range magazine {
		magazineMap[ch]++
	}

	for _, ch := range ransomNote {
		if magazineMap[ch] > 0 {
			magazineMap[ch]--
		} else {
			return false
		}
	}

	return true
}