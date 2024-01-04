package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dyah Lutfia Aziza", "Fia"))
	fmt.Println(strings.Split("Dyah Lutfia Aziza", " "))
	fmt.Println(strings.ToLower("Dyah Lutfia Aziza"))
	fmt.Println(strings.ToUpper("Dyah Lutfia Aziza"))
	fmt.Println(strings.Trim("      Dyah Lutfia Aziza      ", " "))
	fmt.Println(strings.ReplaceAll("Dyah Lutfia Aziza", "Dyah Lutfia Aziza", "Fiadyantha"))
}
