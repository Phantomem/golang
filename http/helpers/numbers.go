package helpers

import "strconv"

func ParseInt(str string) int {
	defaultValue := 0
	val, err := strconv.Atoi(str)
	if err != nil {
		val = defaultValue
	}
	return val
}
