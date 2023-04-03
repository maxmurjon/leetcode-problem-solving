package main

import (
	"fmt"

	countmatching "github.com/leetcode/CountItemsMatchingRule/CountMatching"
)

func main() {
	value := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	fmt.Println(countmatching.CountMatches(value, "color", "silver"))
}
