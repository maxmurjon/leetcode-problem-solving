package funcs

func CheckIfPangram(sentence string) bool {
	var alfabet [26]bool
	for _, v := range sentence {
		if v >= 'a' && v <= 'z' {
			alfabet[v-'a'] = true
		}
	}
	// fmt. Println(alfabet)
	for _, v := range alfabet {
		if !v {
			return false
		}
	}
	return true
}