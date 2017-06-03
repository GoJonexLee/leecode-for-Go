package easy

func findRestaurant(list1 []string, list2 []string) []string {

	f := func(nums []string) map[string]int{
		re := make(map[string]int)
		for k, v := range nums {
			re[v] = k
		}
		return re
	}

	rm := f(list2)
	re, num := []string{}, len(list1)+len(rm)
	for i, v := range list1 {
		if value, ok := rm[v]; ok {
			if num > i + value {
				num = i + value
				re = make([]string, 0)
				re = append(re, v)
			} else if num == i + value {
				re = append(re, v)
			}
		}

	}

	return re
}