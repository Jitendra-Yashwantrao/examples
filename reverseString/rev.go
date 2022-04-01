package reversestring

import (
	"strings"
)

func reverseWords(s string) string {

	slice := strings.Split(s, " ")
	//fmt.Printf("-%v-len-%v-", slice, len(slice))
	reverseString := []string{}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != "" {
			//fmt.Printf("\n-%v-", slice[i])
			reverseString = append(reverseString, slice[i])
		}
	}

	return strings.Join(reverseString, " ")
}
