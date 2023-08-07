package main

// main is the start point
func main() {

	sells := NewSells("123", "987",
		NewProductRepositoryImplementation(),
		NewCustomerRepositoryImplementation())

	sells.GenerateBill(NewEmailSender())
}
