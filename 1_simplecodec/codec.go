package simple

import "strconv"

func ReadNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func WriteNum(i int) string {
	return strconv.Itoa(i)
}
