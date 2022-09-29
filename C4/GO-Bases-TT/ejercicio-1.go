package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	file, err := os.Open("./C4/GO-Bases-TT/customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fscanner := bufio.NewScanner(file)
		for fscanner.Scan() {
			line := fscanner.Text()
			lineSlice := strings.Split(line, "|")
			name := strings.TrimSpace(lineSlice[0])
			salary := strings.TrimSpace(lineSlice[1])
			fmt.Printf("%s %s\n", name, salary)
		}
	}
}
