package rnx

import (
	"encoding/json"
	"encoding/xml"
	"math"
	"testing"
)

var testesCNPJ = []struct {
	n        int64 // input
	expected bool  // expected result
}{
	{79768048000102, true},
	{79768048000103, false},
}

func TestCNPJ(t *testing.T) {
	t.Log("ValidateCNPJ()")
	for _, tt := range testesCNPJ {
		actual := ValidateCNPJ(tt.n)
		if actual != tt.expected {
			t.Errorf("ValidateCNPJ(%d): expected %v, actual %v", tt.n, tt.expected, actual)
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

	pi := math.Pi
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
