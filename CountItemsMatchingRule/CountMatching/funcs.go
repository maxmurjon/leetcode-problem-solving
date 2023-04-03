package countmatching

func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	switch ruleKey {
	case "type":
		for _, v := range items {
			if ruleValue == v[0] {
				count += 1
			}
		}
	case "color":
		for _, v := range items {
			if ruleValue == v[1] {
				count += 1
			}
		}
	case "name":
		for _, v := range items {
			if ruleValue == v[2] {
				count += 1
			}
		}
	}
	return count

}
