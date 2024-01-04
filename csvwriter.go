package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"ibnu", "ahmad", "fauzi"})
	_ = writer.Write([]string{"dyah", "lutfia", "aziza"})
	_ = writer.Write([]string{"satou", "kazuma"})

	writer.Flush()
}
