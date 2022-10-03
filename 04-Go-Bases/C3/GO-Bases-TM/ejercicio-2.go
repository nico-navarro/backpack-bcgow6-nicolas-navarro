package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./C3/GO-Bases-TM/data.csv")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		line := fscanner.Text()
		lineSlice := strings.Split(line, ";")
		id := lineSlice[0]
		price := lineSlice[1]
		quantity := lineSlice[2]
		fmt.Printf("%s\t %s\t %s\n", id, price, quantity)
	}
}
