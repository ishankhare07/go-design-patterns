package factory

import (
	"errors"
	"fmt"
)

// PaymentMethod is the interface for payment methods
type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash = iota
	DebitCard
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

// CashPM implements the PaymentMethod
// interface for cash
type CashPM struct{}

// Pay implements payment for cash type
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// DebitCardPM implements the PaymentMethod
// inteface for Debit card
type DebitCardPM struct{}

// Pay implements payment fot debit card type
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
