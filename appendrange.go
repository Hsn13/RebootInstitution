package piscine

func AppendRange(min, max int) []int {
	var res_arr []int
	if int(max) <= int(min) {
		return res_arr
	}
	for i := min; i < max; i++ {
		res_arr = append(res_arr, i)
	}
	return res_arr
}
