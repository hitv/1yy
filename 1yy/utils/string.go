package utils

import (
	"strconv"
	"strings"
)

func JoinIntsToString(data []int, sep string) string {
	strs := make([]string, len(data))
	for i, v := range data {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, sep)
}
