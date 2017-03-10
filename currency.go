package rnx

import (
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
