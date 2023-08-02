package main

// Account contains attributes of an account entity
type Account struct {
	accountId int
	balance   float64
	nickname  string
}

// NewAccount is the only way to create an account entity
func NewAccount(accountId int) *Account {
	return &Account{
		accountId: accountId,
	}
}
