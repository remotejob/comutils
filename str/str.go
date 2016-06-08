package str

import (
	"unicode"
)

func UpcaseInitial(str string) string {

	runes := []rune(str)

	for i, v := range runes {

		return string(unicode.ToUpper(v)) + string(runes[i+1:])

	}

	return ""

}
