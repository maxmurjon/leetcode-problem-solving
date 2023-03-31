package palindrome

import (
	"strconv"
)

func Palindrome(num int) bool {
	str := strconv.Itoa(num)
	index := 0
	var shart bool = true
	for i := len(str) - 1; i >= len(str)/2; i-- {
		if str[index] == str[i] {
			index += 1
			continue
		} else {
			*&shart = false
			break
		}

	}
	return shart
}
