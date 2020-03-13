package algorithm

// Overlaps 检测两个int数组是否交叉
func Overlaps(list1 []int, list2 []int) bool {
	im := make(map[int]bool)
	for _, item := range list1 {
		im[item] = true
	}

	for _, item := range list2 {
		if im[item] {
			return true
		}
	}

	return false
}
