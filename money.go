package i18n

import (
	"fmt"
	"strconv"

	"github.com/golang-plus/errors"
	"github.com/golang-plus/math/big"
)

// ExchangeableCurrency represents a currency with exchange rate.
type ExchangeableCurrency struct {
	*Currency
	Rate float64
}

// Equal reports two currencies are same.
func (x *ExchangeableCurrency) Equal(y *ExchangeableCurrency) bool {
	return x.Currency.Equal(y.Currency)
}

// NewExchangeableCurrency returns a new exchangeable currency.
func NewExchangeableCurrency(currencyCode string, rate float64) (*ExchangeableCurrency, error) {
	curr, ok := LookupCurrency(currencyCode)
	if !ok {
		return nil, errors.Newf("currency code %q is invalid", currencyCode)
	}
	if big.NewDecimal(rate).Sign() <= 0 {
		return nil, errors.Newf("rate %f is invalid (it should greater than 0)", rate)
	}
	return &ExchangeableCurrency{
		Currency: curr,
		Rate:     rate,
	}, nil
}

// MustNewExchangeableCurrency is similar to NewExchangeableCurrency but panics if error occurred.
func MustNewExchangeableCurrency(currencyCode string, rate float64) *ExchangeableCurrency {
	ec, err := NewExchangeableCurrency(currencyCode, rate)
	if err != nil {
		panic(err)
	}
	return ec
}

// Rounding Mode for Money.
type MoneyRoundingMode byte

// Rounding Mode List.
const (
	MoneyRoundToNearestEven MoneyRoundingMode = iota + 1 // to nearest even (default)
	MoneyRoundToNearestAway                              // to nearest away from zero
	MoneyRoundToZero                                     // to zero
	MoneyRoundAwayFromZero                               // away from zero
)

var (
	DefaultMoneyPrecision    = uint(4)
	DefaultMoneyRoundingMode = MoneyRoundToNearestEven
)

// Money represents an amount (with precision 4) of a specific currency.
type Money struct {
	Currency     *ExchangeableCurrency
	Amount       float64
	RoundingMode MoneyRoundingMode
	Precision    uint // number of decimal digits
}

func (m *Money) round(d *big.Decimal, precision int) *big.Decimal {
	prec := m.Precision
	if precision > 0 {
		prec = uint(precision)
	}
	switch m.RoundingMode {
	case MoneyRoundToNearestEven:
		d.RoundToNearestEven(prec)
	case MoneyRoundToNearestAway:
		d.RoundToNearestAway(prec)
	case MoneyRoundToZero:
		d.RoundToZero(prec)
	case MoneyRoundAwayFromZero:
		d.RoundAwayFromZero(prec)
	default:
		d.RoundToNearestEven(prec)
	}
	return d
}

// Format returns the formatted string of money.
func (m *Money) Format(formatter *CurrencyFormatter) string {
	amount := m.Amount
	if uint(formatter.DecimalDigits) < m.Precision {
		amount, _ = m.round(big.NewDecimal(m.Amount), formatter.DecimalDigits).Float64()
	}
	return formatter.Format(amount)
}

// String returns the formatted string.
// format:  currency code + amount
//			USD 1.2345
func (m *Money) String() string {
	return fmt.Sprintf("%s %."+strconv.Itoa(int(m.Precision))+"f", m.Currency.Code, m.Amount)
}

// Sign returns:
// -1: if m <  0
//  0: if m == 0
// +1: if m >  0
func (m *Money) Sign() int {
	return big.NewDecimal(m.Amount).Sign()
}

// IsZero reports whether the amount is euqal to zero.
func (m *Money) IsZero() bool {
	return m.Sign() == 0
}

// Copy sets x to y, but will not changed if y has changed.
func (x *Money) Copy(y *Money) *Money {
	x.Currency = y.Currency
	x.Amount = y.Amount
	x.Precision = y.Precision
	return x
}

func (m *Money) Exchange(currency *ExchangeableCurrency) *Money {
	if m.Currency.Equal(currency) {
		return m
	}
	sr := big.NewDecimal(m.Currency.Rate)
	tr := big.NewDecimal(currency.Rate)
	amount := m.round(big.NewDecimal(m.Amount).Div(sr).Mul(tr), -1)
	m.Amount, _ = amount.Float64()
	m.Currency = currency
	return m
}

// Add sets amount to the sum of amount and x then returns m.
func (x *Money) Add(y *Money) *Money {
	amount := big.NewDecimal(x.Amount)
	amount.Add(big.NewDecimal(new(Money).Copy(y).Exchange(x.Currency).Amount))
	x.Amount, _ = amount.Float64()
	return x
}

// Sub sets amount to the difference x-y then returns x.
func (x *Money) Sub(y *Money) *Money {
	amount := big.NewDecimal(x.Amount)
	amount.Sub(big.NewDecimal(new(Money).Copy(y).Exchange(x.Currency).Amount))
	x.Amount, _ = amount.Float64()
	return x
}

// Mul sets amount to the product x*y and returns x.
func (x *Money) Mul(y float64) *Money {
	x.Amount, _ = x.round(big.NewDecimal(x.Amount).Mul(big.NewDecimal(y)), -1).Float64()
	return x
}

// Div sets amount to the quotient x/y and return x.
func (x *Money) Div(y float64) *Money {
	x.Amount, _ = x.round(big.NewDecimal(x.Amount).Div(big.NewDecimal(y)), -1).Float64()
	return x
}

// NewMoney returns a new money.
func NewMoney(currency *ExchangeableCurrency, amount float64) (*Money, error) {
	if currency == nil {
		return nil, errors.New("money currency is invalid (it cannot be nil)")
	}
	return &Money{
		Currency:     currency,
		Amount:       amount,
		RoundingMode: DefaultMoneyRoundingMode,
		Precision:    DefaultMoneyPrecision,
	}, nil
}

// MustNewMoney is like as NewMoney but panic if error happens.
func MustNewMoney(currency *ExchangeableCurrency, amount float64) *Money {
	money, err := NewMoney(currency, amount)
	if err != nil {
		panic(err)
	}

	return money
}
