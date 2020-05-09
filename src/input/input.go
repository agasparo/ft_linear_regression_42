package input

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func ReadSTDIN() (float64) {

	fmt.Println("Choose a km")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input := ReplaceWhiteSpace(text)
	f, _ := strconv.ParseFloat(input, 64)
	return (f)
}

func ReplaceWhiteSpace(text string) (string) {

	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\r\n", "")
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\v", "")
	return (text)
}