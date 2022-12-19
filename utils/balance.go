package utils

import (
	"math/big"
)

func GetDisplay(symbol string) string {
	if symbol != "" {
		return symbol
	}
	return "?"
}
func HasBalance(balance *big.Int) bool {
	if balance == nil || (balance != nil && balance.Uint64() == 0) {
		return false
	}
	return true
}
func GetRatioInPool(userBalance, poolTotalSupply *big.Int) *big.Float {
	if userBalance != nil && poolTotalSupply != nil {
		return new(big.Float).Quo(new(big.Float).SetInt(userBalance), new(big.Float).SetInt(poolTotalSupply))
	}
	return new(big.Float)
}
func GetRatioInPoolFloat(userBalance *big.Float, poolTotalSupply *big.Int) *big.Float {
	if userBalance != nil && poolTotalSupply != nil {
		return new(big.Float).Quo(userBalance, new(big.Float).SetInt(poolTotalSupply))
	}
	return new(big.Float)
}

func SupplyBalance(lpAmount, totalSupply, tokenBalance *big.Int) *big.Float {
	tokenRatio := new(big.Float).Quo(BigIntToFloat(tokenBalance), BigIntToFloat(totalSupply))
	return new(big.Float).Mul(BigIntToFloat(lpAmount), tokenRatio)
}

func BigIntToFloat(amount *big.Int) *big.Float {
	return new(big.Float).SetInt(amount)
}
