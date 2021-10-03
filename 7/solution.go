package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	r := ""
	isNegative := false
	s := strconv.FormatInt(int64(x), 10)
	if strings.ContainsRune(s, '-') {
		isNegative = true
		s = s[1:]
	}
	for _, c := range s {
		r = string(c) + r
	}

	result, err := strconv.ParseInt(r, 10, 32)
	if err != nil {
		return 0
	}
	if isNegative {
		return -1 * int(result)
	}
	return int(result)
}

func main() {
	fmt.Println(reverse(-51))
}
