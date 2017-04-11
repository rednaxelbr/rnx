package rnx

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Currency similar ao Delphi
type Currency struct {
	value int64
}

// Value (padrão)
func (c Currency) Value() float64 {
	return float64(c.value) / 10000.0
}

// SetValue (padrão)
func (c *Currency) SetValue(v float64) {
	c.value = Round(v * 10000.0)
}

// String (padrão)
func (c Currency) String() string {
	return fmt.Sprintf("%.4f", c.Value())
}

// MarshalJSON (padrão)
func (c *Currency) MarshalJSON() ([]byte, error) {
	var str string

	str = c.String()
	return []byte(str), nil
}

// UnmarshalJSON (padrão)
func (c *Currency) UnmarshalJSON(curBytes []byte) error {
	s, err := strconv.ParseFloat(string(curBytes), 64)
	if err != nil {
		return err
	}
	c.SetValue(s)
	return nil
}

// NullCurrency represents a Currency that may be null.
// NullCurrency implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullCurrency struct {
	Curr  Currency
	Valid bool // Valid is true if Curr is not NULL
}

// Scan implements the Scanner interface.
func (nc *NullCurrency) Scan(value interface{}) error {
	nc.Curr, nc.Valid = value.(Currency)
	return nil
}

// Value implements the driver Valuer interface.
func (nc NullCurrency) Value() (driver.Value, error) {
	if !nc.Valid {
		return nil, nil
	}
	return nc.Curr, nil
}

// SetValue updates
func (nc NullCurrency) SetValue(v float64) {
	nc.Curr.SetValue(v)
	nc.Valid = true
	return
}
