package rnx

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

var testesRound = []struct {
	n        float64 // input
	expected int64   // expected result
}{
	{1.02, 1},
	{2.45, 2},
	{3.51, 4},
	{4.99, 5},
	{5.00, 5},
	{-4.3, -4},
}

func TestRound(t *testing.T) {
	t.Log("Round()")
	for _, tt := range testesRound {
		actual := Round(tt.n)
		if actual != tt.expected {
			t.Errorf("Round(%.4f): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestCurrency(t *testing.T) {
	var valor Currency
	type testeJSON struct {
		Campo Currency `json:"campo"`
	}
	type testeXML struct {
		XMLName xml.Name `xml:"teste"`
		Campo   Currency `xml:"campo"`
	}

	var pi float64
	pi = 355.0 / 113.0 //aproximação de Pi
	t.Log("Pi():", pi)
	expected := "3.1416"
	valor.SetValue(pi)
	actual := valor.String()
	if actual != expected {
		t.Errorf("Currency.String(%.4f): expected %s, actual %s", valor.Value(), expected, actual)
	}

	t.Log("JSON")
	expected = "{\"campo\":3.1416}"
	testej := testeJSON{Campo: valor}
	b, _ := json.Marshal(testej)
	actual = string(b)
	if actual != expected {
		t.Errorf("Currency.MarshalJSON(%.4f): expected %s, actual %s", valor.Value(), expected, actual)
	}

	t.Log("XML")
	expected = "<teste><campo>3.1416</campo></teste>"
	testex := testeXML{Campo: valor}
	b, _ = xml.Marshal(testex)
	actual = string(b)
	if actual != expected {
		t.Errorf("Currency.MarshalXML(%.4f): expected %s, actual %s", valor.Value(), expected, actual)
	}

}
