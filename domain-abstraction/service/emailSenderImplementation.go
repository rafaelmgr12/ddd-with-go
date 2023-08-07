package main

import "fmt"

// EmailSenderImplementation implements EmailSender interface
type EmailSenderImplementation struct {
}

func NewEmailSender() EmailSender {
	return &EmailSenderImplementation{}
}

func (esi *EmailSenderImplementation) SendEmail(customer *Customer, product *Product) {
	fmt.Printf("Invoking SMTP server \n")
	fmt.Printf("Sending email to %s. Product bought is %s \n", customer.emailAddress, product.name)
}
