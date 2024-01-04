package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "ibnu.ahmad.fauzi\n" +
		"dyah.lutfia.aziza\n" + "satou.kazuma"

	reader := csv.NewReader(strings.NewReader((csvString)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
