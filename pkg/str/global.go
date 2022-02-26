package str

import "strconv"

//StringToInt convert string to int
func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}
