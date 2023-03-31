package flippingimage

func FlipAndInvertImage(image [][]int)[][]int{
	newSlice:=[][]int{}
	for _,v:=range image{
		slice1:=[]int{}
		for i := len(v)-1; i >=0; i-- {
			if v[i]==0{
				slice1 = append(slice1, 1)
			}else if v[i]==1{
				slice1=append(slice1, 0)
			}
		}
		newSlice = append(newSlice, slice1)
	}
	return newSlice
}