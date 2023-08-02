package main

import "time"

// Transaction contains attributes of transaction aggregate
type Transaction struct {
	id       string
	bankName string
	customer *Customer
	amount   *Amount
	transfer *Transfer
}

// NewTransaction is the only way to create a transaction object and its entities and value objects
func NewTransaction(id string, bankName string, amount float64, currency string, date time.Time, accountId int, customerId int) *Transaction {
	return &Transaction{
		id:       id,
		bankName: bankName,
		customer: NewCustomer(customerId),
		amount:   NewAmount(amount, currency),
		transfer: NewTransfer(id, date, accountId),
	}
}

// TransferMoney validates date and limit before executing a transfer.
// Date and limit validation is based on transfer and amount artifacts
func (t *Transaction) TransferMoney() bool {

	if !t.transfer.IsDateAllowed() {
		return false
	}

	if !t.amount.IsAmountLimitAllowed() {
		return false
	}

	return t.executeTransfer()
}

// executeTransfer is an unexported method that executes a transfer as such
func (t *Transaction) executeTransfer() bool {
	// here is where transfer logic should be placed
	return true
}
