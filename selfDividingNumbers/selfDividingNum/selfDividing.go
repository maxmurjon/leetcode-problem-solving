package selfdividingnum

func SelfDividingNumbers(left int, right int) []int {
	arr := []int{}
	for i := left; i <= right; i++ {
		if isDividable(i) {
			arr = append(arr, i)
		}
	}
	return arr
}
func isDividable(i int) bool {
	num:=i
	for num > 0 {
		digit := num % 10
		if digit == 0 || i%digit != 0 {
			return false
		}
		num/= 10

	}
	return true
}