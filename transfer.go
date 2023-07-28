package main

import "time"

// Transfer contains attributes of a transfer aggregate
type Transfer struct {
	id      string
	date    time.Time
	account *Account
}

// NewTransfer is the only way to create a transfer aggregate
func NewTransfer(id string, date time.Time, accountId int) *Transfer {

	return &Transfer{
		id:      id,
		date:    date,
		account: NewAccount(accountId),
	}
}

// IsDateAllowed checks whether transfer is executed in the current day
func (t *Transfer) IsDateAllowed() bool {
	now := time.Now()
	if t.date.After(now) {
		return false
	}
	return true
}
