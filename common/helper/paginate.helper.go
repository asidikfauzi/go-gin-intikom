package helper

func GetOffset(pageParam, limitParam int) int {

	if pageParam == 0 {
		pageParam = 1
	}

	offset := (pageParam - 1) * limitParam

	return offset
}
