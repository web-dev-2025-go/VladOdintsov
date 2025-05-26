package utils

import "strings"

func FormatSlice(slice []string) string {
	return strings.Join(slice, ",")
}
