package camelcase_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	res := CamelCase("cats AND*Dogs-are Awesome")
	fmt.Println(res)
}

func CamelCase(str string) string {

	for i := 0; i < len(str); i++ {
		s := string(str[i])

		if i == 0 {
			s = strings.ToLower(s)
			continue
		}

		if str[i-1] >= 32 && str[i-1] <= 47 {
			s = strings.ToUpper(s)

			s = strings.Replace("", "", "", int(str[i]))

			i--
		} else {
			s = strings.ToLower(s)
		}

	}

	return str

}
