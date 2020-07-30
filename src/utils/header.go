package utils

import (
	"fmt"
	"strings"
)

func PrintHeader() {
	PrintBorder()
	fmt.Printf("%-10s %-50s %-20s\n", "Order", "Item", "Serving Time")
	PrintBorder()

}
func PrintBorder() {
	fmt.Printf("%120s\n", strings.Repeat("=", 120))
}

func Trimming(s string) string {
	return strings.Trim(s, " ")
}
