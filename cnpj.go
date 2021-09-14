package rnx

import "fmt"

//=============================================================================

// ValidateCNPJ valida um CNPJ
func ValidateCNPJ(cnpj int64) bool {
	var result bool

	if cnpj > 99999999999999 {
		return false
	}
	strCNPJ := fmt.Sprintf("%014d", cnpj)
	semDV := strCNPJ[0:12]

	//--- primeiro dígito
	mult := 5
	soma := 0
	for i := range semDV {
		soma += int(semDV[i]-'0') * mult
		mult--
		if mult < 2 {
			mult = 9
		}
	}
	d1 := 11 - (soma % 11)
	if d1 >= 10 {
		d1 = 0
	}

	//--- segundo dígito
	mult = 6
	soma = 0
	for i := range semDV {
		soma += int(semDV[i]-'0') * mult
		mult--
		if mult < 2 {
			mult = 9
		}
	}
	soma += d1 * mult
	d2 := 11 - (soma % 11)
	if d2 >= 10 {
		d2 = 0
	}

	dv1 := int(strCNPJ[12] - '0')
	dv2 := int(strCNPJ[13] - '0')
	result = (d1 == dv1 && d2 == dv2)

	return result
}
