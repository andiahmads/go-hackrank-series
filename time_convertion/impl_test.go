package timeconvertion

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTimeConverstion(t *testing.T) {

	// TimeConvertion(07, 6, 9, "PM")

	res := TimeConvertionTwo("07:05:45PM")
	fmt.Println(res)

}

func TimeConvertion(h, m, s int, suffix string) {

	if suffix == "AM" && h == 12 {
		h = 0
	}
	if suffix == "PM" && h != 12 {
		h += 12
	}
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}

func TimeConvertionTwo(s string) string {
	s = strings.ToLower(s)

	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")

	t := s[:len(s)-2]

	timeArr := strings.Split(t, ":")
	fmt.Println(timeArr)

	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]

	hhInt, err := strconv.Atoi(hh) // string to int

	if err != nil {
		panic(err.Error())
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt)

	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)

}
