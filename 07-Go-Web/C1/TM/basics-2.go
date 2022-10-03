package main

import (
	"fmt"
	"net/http"
)

func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola\n")                   //response al cliente
	fmt.Println("Alguien accedió el endpoint") //esto ocurre en el sv
}

func main() {
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}
