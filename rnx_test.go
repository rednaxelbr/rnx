package rnx

import "testing"

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
	for _, tt := range testesRound {
		actual := Round(tt.n)
		if actual != tt.expected {
			t.Errorf("Round(%.4f): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func TestCurrency(t *testing.T) {
	var valor Currency
	expected := "3.1416"
	valor.SetValue(355.0 / 113.0) //aproximação de Pi = 3.14159292
	actual := valor.String()
	if actual != expected {
		t.Errorf("Currency.String(%.4f): expected %s, actual %s", valor.Value(), expected, actual)
	}
}
