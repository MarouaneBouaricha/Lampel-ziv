package main

import (
	"fmt"
	"log"
)

type Tuple struct {
	key  int
	char string
}

func zip(keys []int, chars []string) ([]Tuple, error) {

	if len(keys) != len(chars) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	output := make([]Tuple, len(keys), len(chars))

	for i, elt := range keys {
		output[i] = Tuple{elt, chars[i]}
	}

	return output, nil
}

func compress(msg string) []Tuple {
	dict := map[string]int{}
	prefix := ""
	dict_index := 1
	charTab, KeyTab := []string{}, []int{}
	var code_prefix int
	for _, v := range msg {
		if _, ok := dict[prefix+string(v)]; ok {
			prefix += string(v)
		} else {
			if prefix == "" {
				code_prefix = 0
			} else {
				code_prefix = dict[prefix]
			}
			charTab = append(charTab, string(v))
			KeyTab = append(KeyTab, code_prefix)
			dict[prefix+string(v)] = dict_index
			dict_index += 1
			prefix = ""
		}
	}
	if prefix != "" {
		code_prefix = dict[prefix]
		charTab = append(charTab, "")
		KeyTab = append(KeyTab, code_prefix)
	}

	output, err := zip(KeyTab, charTab)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func decompress(tuple []Tuple) string {
	dict := map[int]string{}
	str := ""
	output := ""
	dict_index := 1
	var code int
	var char string
	for _, elt := range tuple {
		code = elt.key
		char = elt.char
		if code == 0 {
			str = ""
		} else {
			str = dict[code]
		}
		output += str + char
		dict[dict_index] = str + char
		dict_index += 1

	}
	return output
}

func main() {
	msg := "ABBCBCABABCAABCAAB"
	fmt.Println(compress(msg))
	h := compress(msg)
	fmt.Println(decompress(h))

}
