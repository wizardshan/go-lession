package response

import (
	"fmt"
	"github.com/shopspring/decimal"
)
type Money int

func (resp Money) MarshalJSON() ([]byte, error) {
	d, _ := decimal.NewFromFloat(float64(resp)).Mul(decimal.NewFromFloat(0.01)).Float64()
	str := fmt.Sprintf("%.2f", d)
	return []byte(str), nil
}