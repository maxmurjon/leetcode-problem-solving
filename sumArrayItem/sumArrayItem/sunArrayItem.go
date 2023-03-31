package sumarrayitem


func RunningSum(nums []int) []int {
	resultArray:=[]int{}
	for i:=range nums{
		result:=0
		for j := 0; j <= i; j++ {
			result+=nums[j]
		}
		resultArray = append(resultArray, result)
	}
	return resultArray
}