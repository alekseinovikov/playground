package internal

import "fmt"

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64 | int
}

func toString[N Number](num N) string {
	return fmt.Sprint(num)
}

func Sum[N Number](params ...N) N {
	var result N = 0.0
	for _, val := range params {
		result = result + val
	}

	return result
}

func ConcatStringAndNumber[N Number](str string, num N) string {
	return str + toString(num)
}
