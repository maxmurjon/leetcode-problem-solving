package sumfunc

func TwoSum(nums []int,target int)(result []interface{}){
	dict:=make(map[int]int)
	for i,v:=range nums{
		s:=target-v
		if dict[s]!=0{
			result = append(result,dict[s]-1,i)
			return
		}
		dict[v]=i+1
	}
	return
}