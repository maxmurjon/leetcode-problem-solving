package palindromearray

func FirstPalindrome(words []string) string {
	for _,v:=range words{
		if v==reverse(v){
			return v
		}else{
			continue
		}
	}
	return ""
}

func reverse(str string)string{
	strRet:=""
	for i := len(str)-1; i>=0; i-- {
		strRet+=string(str[i])
	}
	return strRet
}