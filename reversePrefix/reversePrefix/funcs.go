package reverseprefix

func ReversePrefix(word string, ch byte) string {
	str := ""
	ifs:=false
	for _, v := range word {
		if v == rune(ch) {
			str += string(v)
			*&ifs=true
			break
		}
		str += string(v)
	}
	newStr := ""
	if !ifs{
		return word
	}
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	newStr += word[len(newStr):]
	return newStr
}
