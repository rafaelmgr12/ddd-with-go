package main

// EmailSender defines functionality to send an email
type EmailSender interface {

	// SendEmail receives customer and product to send an email
	SendEmail(*Customer, *Product)
}
