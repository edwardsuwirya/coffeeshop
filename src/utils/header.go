package utils

import (
	"fmt"
	"strings"
)

func PrintHeader() {
	PrintBorder()
	fmt.Printf(FORMAT_STRING_HEADER, "Order", "Item", "Serving Time")
	PrintBorder()

}
func PrintBorder() {
	fmt.Printf(FORMAT_STRING_BORDER, strings.Repeat("=", BORDER_LENGTH))
}

func Trimming(s string) string {
	return strings.Trim(s, " ")
}
