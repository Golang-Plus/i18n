package i18n

import (
	"fmt"

	"github.com/golang-plus/errors"
	"github.com/golang-plus/math/big"
)

// Money represents an amount (with precision 4) of a specific currency.
type Money struct {
	amount   *big.Decimal
	Currency *Currency
	Amount   float64
}

// Format returns the formatted string of money.
func (m *Money) Format(formatter *CurrencyFormatter) string {
	return formatter.Format(m.Amount)
}

// String returns the formatted string (format: currency code + amount).
func (m *Money) String() string {
	return fmt.Sprintf("%s %f", m.Currency.Code, m.Amount)
}

// IsZero reports whether the amount is euqal to zero.
func (m *Money) IsZero() bool {
	return m.amount.IsZero()
}

// Sign returns:
// -1: if m <  0
//  0: if m == 0
// +1: if m >  0
func (m *Money) Sign() int {
	return m.amount.Sign()
}

// Add sets amount to the sum of amount and x then returns m.
func (m *Money) Add(x float64) *Money {
	m.amount.Add(big.NewDecimal(x))
	m.Amount, _ = m.amount.Float64()
	return m
}

// Sub sets amount to the difference amount-x then returns m.
func (m *Money) Sub(x float64) *Money {
	m.amount.Sub(big.NewDecimal(x))
	m.Amount, _ = m.amount.Float64()
	return m
}

// Mul sets amount to the product amount*x and returns m.
func (m *Money) Mul(x float64) *Money {
	m.amount.Mul(big.NewDecimal(x))
	m.Amount, _ = m.amount.Float64()
	return m
}

// Div sets amount to the quotient amount/x and return m.
func (m *Money) Div(x float64) *Money {
	m.amount.Div(big.NewDecimal(x))
	m.Amount, _ = m.amount.Float64()
	return m
}

// RoundToNearestEven rounds the amount by "To Nearest Even" mode with 4 precision.
// precision indicates the number of digits after the decimal point.
func (m *Money) RoundToNearestEven(precision uint) *Money {
	m.amount.RoundToNearestEven(precision)
	m.Amount, _ = m.amount.Float64()
	return m
}

// RoundToNearestEven rounds the amount by "To Nearest Away From Zero" mode with 4 precision.
// precision indicates the number of digits after the decimal point.
func (m *Money) RoundToNearestAway(precision uint) *Money {
	m.amount.RoundToNearestAway(precision)
	m.Amount, _ = m.amount.Float64()
	return m
}

// RoundToZero rounds the amount by "To Zero" mode with 4 precision.
// precision indicates the number of digits after the decimal point.
func (m *Money) RoundToZero(precision uint) *Money {
	m.amount.RoundToZero(precision)
	m.Amount, _ = m.amount.Float64()
	return m
}

// Round is same as RoundToNearestEven.
func (m *Money) Round(precision uint) *Money {
	return m.RoundToNearestEven(precision)
}

// Truncate is same as RoundToZero.
func (m *Money) Truncate(precision uint) *Money {
	return m.RoundToZero(precision)
}

// NewMoney returns a new money.
func NewMoney(currency *Currency, amount float64) (*Money, error) {
	if currency == nil {
		return nil, errors.New("currency of money cannot be nil")
	}

	return &Money{
		amount:   big.NewDecimal(amount),
		Currency: currency,
		Amount:   amount,
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
