package main

import (
	"fmt"

	flippingimage "github.com/leetcode/flippingImage/flippingImage"
)

func main(){
	fmt.Println(flippingimage.FlipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}}))
}