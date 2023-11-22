package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// AccountBalance represents a single row inside the "accound_denom_balancee" table
type AccountBalance struct {
	Account string `db:"account"`
	Denom   string `db:"denom"`
	Aalance string `db:"amount"`
}

// NewSupplyRow allows to easily create a new NewSupplyRow
func NewAccountBalance(account string, coin sdk.Coin) AccountBalance {
	return AccountBalance{
		Account: account,
		Denom:   coin.Denom,
		Aalance: coin.Amount.String(),
	}
}

// Equals return true if one totalSupplyRow representing the same row as the original one
func (v AccountBalance) Equals(w AccountBalance) bool {
	return v == w
}
