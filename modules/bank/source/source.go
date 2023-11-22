package source

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/forbole/bdjuno/v4/types"
)

type Source interface {
	GetBalances(addresses []string, height int64) ([]types.AccountBalance, error)
	GetSupply(height int64) (sdk.Coins, error)
	GetAccountDenomBalance(address string, denom string) (*sdk.Coin, error)

	// -- For hasura action --
	GetAccountBalance(address string, height int64) ([]sdk.Coin, error)
}
