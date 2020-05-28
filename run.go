package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseMap(m map[rune]rune) map[rune]rune {
	n := make(map[rune]rune)
	for k, v := range m {
		n[v] = k
	}
	return n
}

var dict = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5',
}
var reversed_dict = reverseMap(dict)

func coder(dict map[rune]rune) func(string) string {

	return func(in string) string {
		out := strings.Map(
			func(r rune) rune {
				ch, ok := dict[unicode.ToLower(r)]
				if ok {
					return ch
				}
				return r
			},
			in)
		return out
	}

}

var encode = coder(dict)
var decode = coder(reversed_dict)



func main() {

	input := "hello"
	fmt.Println(encode(input))
	fmt.Println(decode(encode(input)))

}
