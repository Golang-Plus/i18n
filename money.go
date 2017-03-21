package i18n

import (
	"fmt"

	"github.com/golang-plus/errors"
	"github.com/golang-plus/math/big"
)

// Money represents an amount of a specific currency.
type Money struct {
	Currency *Currency
	Amount   float64
	decimal  *big.Decimal
}

// NewMoney returns a new money.
func NewMoney(currency *Currency, amount float64) (*Money, error) {
	if currency == nil {
		return nil, errors.New("currency of money cannot be nil")
	}

	return &Money{
		Currency: currency,
		Amount:   amount,
		decimal:  big.NewDecimal(amount),
	}, nil
}

// MustNewMoney is like as NewMoney but panic if error happens.
func MustNewMoney(currency *Currency, amount float64) *Money {
	money, err := NewMoney(currency, amount)
	if err != nil {
		panic(err)
	}

	return money
}

// Format returns the formatted string of money.
func (m *Money) Format(formatter *CurrencyFormatter) string {
	return formatter.Format(m.Amount)
}

// String returns the string of money (format: currency code + amount).
func (m *Money) String() string {
	return fmt.Sprintf("%s %f", m.Currency.Code, m.decimal)
}

// Add sets m.decimal to the sum of m.decimal and another and returns m.
func (m *Money) Add(x float64) *Money {
	m.decimal.Add(big.NewDecimal(x))
	m.Amount = m.decimal.Float64()
	return m
}

// Sub sets m.decimal to the difference m.decimal-another and returns m.
func (m *Money) Sub(x float64) *Money {
	m.decimal.Sub(big.NewDecimal(x))
	m.Amount = m.decimal.Float64()
	return m
}

// Mul sets m.decimal to the product m.decimal*another and returns m.
func (m *Money) Mul(x float64) *Money {
	m.decimal.Mul(big.NewDecimal(x))
	m.Amount = m.decimal.Float64()
	return m
}

// Div sets m.decimal to the quotient m/another and return m.
func (m *Money) Div(x float64) *Money {
	m.decimal.Div(big.NewDecimal(x))
	m.Amount = m.decimal.Float64()
	return m
}
