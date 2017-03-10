// Package rnx - Biblioteca Rednaxel
package rnx

// round porque o Golang não tem...
func Round(val float64) int64 {
    if val < 0 {
        return int64(val - 0.5)
    }
    return int64(val + 0.5)
}

