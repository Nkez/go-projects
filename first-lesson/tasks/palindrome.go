package tasks

func CheckPalindrome(word string) bool {
	if len(word) < 2 {
		return true
	}

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}

	return true
}
