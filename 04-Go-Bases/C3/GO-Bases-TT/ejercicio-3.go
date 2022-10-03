package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Service struct {
	Name    string
	Price   float64
	Minutes int
}

type Maintenance struct {
	Name  string
	Price float64
}

func addProducts(products []Product, totalChannel chan float64) {
	var totalPrice float64
	for _, product := range products {
		totalPrice += product.Price * float64(product.Quantity)
	}
	totalChannel <- totalPrice
}

func addServices(services []Service, totalChannel chan float64) {
	var totalPrice float64
	for _, service := range services {
		minutesCharged := service.Minutes
		if service.Minutes < 30 {
			minutesCharged = 30
		}
		totalPrice += service.Price * (float64(minutesCharged) / 30)
	}
	totalChannel <- totalPrice
}

func addMaintenances(maintenances []Maintenance, totalChannel chan float64) {
	var totalPrice float64
	for _, maintenance := range maintenances {
		totalPrice += maintenance.Price
	}
	totalChannel <- totalPrice
}

func main() {
	product1 := Product{"Nike", 250, 2}
	service1 := Service{"Servicio", 2500, 25}
	maintenance1 := Maintenance{"Mantenimiento", 3000}

	product2 := Product{"Nike", 250, 2}
	service2 := Service{"Servicio", 2500, 25}
	maintenance2 := Maintenance{"Mantenimiento", 3000}

	products := []Product{product1, product2}
	services := []Service{service1, service2}
	maintenances := []Maintenance{maintenance1, maintenance2}

	channelProducts := make(chan float64)
	channelServices := make(chan float64)
	channelMaintenances := make(chan float64)

	go addProducts(products, channelProducts)
	go addServices(services, channelServices)
	go addMaintenances(maintenances, channelMaintenances)

	totalPriceProducts := <-channelProducts
	totalPriceServices := <-channelServices
	totalPriceMaintenances := <-channelMaintenances

	totalPrice := totalPriceProducts + totalPriceServices + totalPriceMaintenances
	fmt.Printf("El precio total de todo es: %f\n", totalPrice)
}
