package ethex

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

// FromWei convert from wei
func FromWei(v *big.Int, d int) decimal.Decimal {
	return decimal.NewFromBigInt(v, 0).DivRound(decimal.NewFromFloat(math.Pow10(d)), 32)
}
