package main

import "fmt"

// main is the start point
func main() {

	sells := NewSells("123", "987",
		NewProductRepositoryImplementation(),
		NewCustomerRepositoryImplementation())

	fmt.Println("Information about a bill: ")
	sells.GenerateBill()
}
