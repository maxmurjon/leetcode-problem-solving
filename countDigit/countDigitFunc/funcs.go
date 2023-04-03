package countdigitfunc

import "strconv"




func CountDigits(num int) int {
	// test:=make(map[string]bool)
	count:=0
	for _,v:=range strconv.Itoa(num){
		n,_:=strconv.Atoi(string(v))
		if num%n==0{
				count+=1
		}	
	}
	return count
}

