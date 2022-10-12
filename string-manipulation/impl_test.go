package stringmanipulation_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()

		toUpper := strings.ToUpper(line)
		fmt.Println(toUpper)
	}

}
