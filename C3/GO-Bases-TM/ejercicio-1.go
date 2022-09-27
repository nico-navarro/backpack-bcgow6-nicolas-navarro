package main

import (
	"fmt"
	"os"
)

func main() {
	d1 := []byte("ID;Precio;Cantidad \n1;3500;12 \n2;1000;1 \n3;3200;4")
	err := os.WriteFile("./C3/GO-Bases-TM/data.csv", d1, 0777)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Todo ok\n")
	}
}
