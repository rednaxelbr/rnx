// Package rnx Copyright(C) 2017 Rednaxel Informática Ltda
package rnx

import (
	"bytes"
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

const (
	currPrecision = 10000.0 // 4 decimals
)

//=============================================================================

// Currency similar ao Delphi
type Currency struct {
	value int64
}

// Value (padrão)
func (c Currency) Value() float64 {
	return float64(c.value) / currPrecision
}

// SetValue (padrão)
func (c *Currency) SetValue(v float64) {
	c.value = Round(v * currPrecision)
}

// String (nunca usar asterisco)
func (c Currency) String() string {
	return fmt.Sprintf("%.4f", c.Value())
}

// MarshalJSON (nunca usar asterisco)
func (c Currency) MarshalJSON() ([]byte, error) {
	str := c.String()
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

// MarshalXML (padrão)
func (c Currency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	v := c.String()
	e.EncodeElement(&v, start)
	return nil
}

// UnmarshalXML (padrão)
func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	s, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	c.SetValue(s)
	return nil
}

//=============================================================================

// NullCurrency represents a Currency that may be null.
// NullCurrency implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullCurrency struct {
	Curr  Currency
	Valid bool // Valid is true if Curr is not NULL
}

func (nc *NullCurrency) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}

// Scan implements the Scanner interface.
func (nc *NullCurrency) Scan(value interface{}) error {
	switch value.(type) {
	case float32, float64:
		nc.Curr.SetValue(value.(float64))
		nc.Valid = true
	case []uint8:
		str := string(nc.getBytes(value))
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		nc.Curr.SetValue(v)
		nc.Valid = true
	case nil:
		return nil
	default:
		return fmt.Errorf("Tipo inválido: %s", reflect.TypeOf(value))
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nc NullCurrency) Value() (driver.Value, error) {
	if !nc.Valid {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%.4f", nc.Curr.Value())
	return buf.Bytes(), nil
}

// SetValue updates
func (nc *NullCurrency) SetValue(v float64) {
	nc.Curr.SetValue(v)
	nc.Valid = true
	return
}

// String (nunca usar asterisco)
func (nc NullCurrency) String() string {
	if nc.Valid {
		return fmt.Sprintf("%.4f", nc.Curr.Value())
	}
	return ""
}

// MarshalXML (padrão)
func (nc NullCurrency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type str struct{ val string }
	v := &str{nc.String()}
	e.Encode(v)
	return nil
}

// MarshalJSON (padrão)
func (nc NullCurrency) MarshalJSON() ([]byte, error) {
	if nc.Valid {
		str := nc.String()
		return []byte(str), nil
	}
	return []byte("null"), nil
}

// UnmarshalJSON (padrão)
func (nc *NullCurrency) UnmarshalJSON(curBytes []byte) error {
	if curBytes == nil {
		nc.SetValue(0)
		nc.Valid = false
		return nil
	}
	//log.Printf("NullCurrency.UnmarshalJSON(%v)", curBytes)
	if string(curBytes) != "null" {
		s, err := strconv.ParseFloat(string(curBytes), 64)
		if err != nil {
			return err
		}
		nc.SetValue(s)
	} else {
		nc.SetValue(0)
		nc.Valid = false
	}
	return nil
}

//=============================================================================

// NewCurrency creates a new Currency
func NewCurrency(v float64) Currency {
	return Currency{
		value: Round(v * currPrecision)}
}
