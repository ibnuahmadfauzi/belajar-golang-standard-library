package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Ibnu", "Ahmad", "Fauzi", "Dyah", "Lutfia", "Aziza"}
	values := []int{100, 70, 60, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Dyah"))
	fmt.Println(slices.Index(names, "Dyah"))
	fmt.Println(slices.Index(names, "Megumin"))
}
