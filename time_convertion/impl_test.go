package timeconvertion

import (
	"fmt"
	"testing"
)

func TestTimeConverstion(t *testing.T) {

	TimeConvertion(07, 6, 9, "PM")

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
