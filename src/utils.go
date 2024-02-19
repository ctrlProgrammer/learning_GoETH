package LETH

import (
	"math"
	"math/big"
)

func ParseETH(wei big.Int) string {
	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return string(ethValue.String())
}
