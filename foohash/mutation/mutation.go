package mutation

import "fmt"
import "strings"

var leetCharset = map[string]string{
	"o": "0",
	"l": "1",
	"e": "3",
	"s": "5",
	"t": "7",
}

func Leet(src string) (dst string) {
	for char, leet := range leetCharset {
		dst = strings.Replace(src, char, leet, -1)
	}
	return dst
}

func NumSuffixes(src string) (dsts []string) {
	var commons = []int{123, 1234, 12345, 123456, 1234567, 12345678, 123456789, 1234567890}
	for _, i := range commons {
		dst := fmt.Sprintf("%s%d", src, i)
		dsts = append(dsts, dst)
	}
	for i := 0; i <= 100; i++ {
		dst := fmt.Sprintf("%s%d", src, i)
		dsts = append(dsts, dst)
	}
	for i := 1900; i <= 2050; i++ {
		dst := fmt.Sprintf("%s%d", src, i)
		dsts = append(dsts, dst)
	}
	return dsts
}
