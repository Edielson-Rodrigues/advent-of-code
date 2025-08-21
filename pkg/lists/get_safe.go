package lists

func GetSafe(list []int, index int) int {
	if len(list) > index {
		return list[index]
	}

	return 0
}
