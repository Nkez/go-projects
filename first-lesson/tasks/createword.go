package tasks

func CreateWord(word string) string {

	array := make([]rune, len(word))
	emptyArray := make([]rune, len(word))
	for _, c := range word {
		if c == ')' {
			i := len(array) - 1
			emptyArray = emptyArray[:0]
			for ; array[i] != '('; i-- {
				emptyArray = append(emptyArray, array[i])
			}
			array = append(array[:i], emptyArray...)
		} else {
			array = append(array, c)
		}
	}

	return string(array)
}
