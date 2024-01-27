package v2

func MemoryWaste(n int) []int {
	var data = make([]int, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, i)
	}
	return data
}
