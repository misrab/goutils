package utils

import (
	"testing"
	//"fmt"
)

 // ReadCsv(path string, c chan []string, screenBlank bool, delimiter rune)
func TestReadCsv(t *testing.T) {
	path := "data/csvtest.csv"
	c := make(chan []string)

	go ReadCsv(path, c, true, ',')

	// for record := range c {
	// 	fmt.Println(record)
	// }
}