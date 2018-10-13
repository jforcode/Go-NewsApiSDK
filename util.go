package newsApi

func minInt(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func isEmptyArr(arr []string) bool {
	return arr == nil || len(arr) == 0
}
