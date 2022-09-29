package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Customer struct {
	ID      int //Legajo
	Name    string
	DNI     string
	Number  int
	Address string
}

var lastID = 0

func generateID() int {
	lastID++
	return lastID
}

func customerExists(newDNI string) error {
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
			DNI := strings.TrimSpace(lineSlice[2])
			if newDNI == DNI {
				return errors.New("ya existe")
			}
		}
	}
	file.Close()
	return err
}

func validFields(name string, DNI string, number int, address string) (err error) {
	if name == "" || DNI == "" || number == 0 || address == "" {
		err = errors.New("se ingresó al menos un valor cero")
		return
	}
	return
}

func newCustomer(name string, DNI string, number int, address string) (customer Customer, err error) {
	err = customerExists(DNI)
	if err != nil {
		return
	}
	err = validFields(name, DNI, number, address)
	if err != nil {
		return
	}
	customer.ID = generateID()
	customer.Name = name
	customer.DNI = DNI
	customer.Number = number
	customer.Address = address
	return
}

func main() {
	_, err := newCustomer("Nico", "1234K", 1234, "AD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fin de la ejecución")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println("No han quedado archivos abiertos")
}
