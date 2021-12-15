package request

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
)
type Money struct {
	Val int
}

func (req *Money) UnmarshalJSON(b []byte) error {
	float, err := strconv.ParseFloat(string(b),64)
	if err != nil {
		return err
	}

	if float <= 0 {
		return errors.New("金额不能为负数")
	}

	req.Val = int(decimal.NewFromFloat(float).Mul(decimal.NewFromFloat(100)).IntPart())
	return nil
}

func (req *Money) Value() int {
	return req.Val
}

