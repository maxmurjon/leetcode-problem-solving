package commonfactor

func CommonFactors(a int, b int) int {
	var count int
	if a > b {
		for i := 1; i <= b; i++ {
			if a%i == 0 && b%i == 0 {
				count += 1
			}
		}
	} else {
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				count += 1
			}
		}
	}
	return count
}
