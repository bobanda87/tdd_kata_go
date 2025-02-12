package string_calculator

import (
	"strconv"
	"strings"
)

func Add(expr string) int64 {
	if expr == "" {
		return 0
	}

	delimiter := ","
	if strings.HasPrefix(expr, "//") {
		delimiter_start := 2
		if expr[delimiter_start] == '[' {
			delimiter_end := strings.Index(expr, "]")
			delimiter = string(expr[delimiter_start+1 : delimiter_end])
			expr = strings.Replace(expr, expr[0:delimiter_end+1], "", -1)
		} else {
			delimiter = string(expr[2])
			expr = strings.Replace(expr, "//"+delimiter, "", -1)
		}
	}

	numbers := strings.Split(strings.Replace(expr, "\n", delimiter, -1), delimiter)
	sum := int64(0)
	for i := 0; i < len(numbers); i++ {
		i, _ := strconv.ParseInt(numbers[i], 10, 64)
		sum += i
	}

	return sum
}
