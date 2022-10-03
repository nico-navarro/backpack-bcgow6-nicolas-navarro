package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type product struct {
		Name      string
		Price     int
		Published bool
	}
	// p := product{
	// 	Name:      "MacBook Pro",
	// 	Price:     1500,
	// 	Published: true,
	// }
	// jsonData, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", jsonData)

	jsonData := []byte(`{"Name": "MacBook Air", "Price": 900, "Published": true}`)
	var p product
	if err := json.Unmarshal(jsonData, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v", p)
}
