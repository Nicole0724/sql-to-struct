package words

import (
	"regexp"
	"strings"
)

// ToBigCamelCase 大驼峰
func ToBigCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, w := range words {
		words[i] = strings.Title(w)
	}
	return strings.Join(words, "")
}

// ToCamelCase 小驼峰
func ToCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, w := range words[1:] {
		words[i+1] = strings.Title(w)
	}
	return strings.Join(words, "")
}
