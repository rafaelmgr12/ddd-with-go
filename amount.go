package main

// Amount contains attributes of an amount value object
type Amount struct {
	value    float64
	currency string
}

// NewAmount is the only way to create an amount value object
func NewAmount(value float64, currency string) *Amount {
	return &Amount{
		value:    value,
		currency: currency,
	}
}

// IsAmountLimitAllowed check whether amount to transfer does not got beyond the bank limits
func (a *Amount) IsAmountLimitAllowed() bool {
	if a.value >= 10000 {
		return false
	}
	return true
}
