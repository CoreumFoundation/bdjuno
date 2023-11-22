package database

import (
	"fmt"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lib/pq"
)

// SaveSupply allows to save for the given height the given total amount of coins
func (db *Db) SaveSupply(coins sdk.Coins, height int64) error {
	query := `
INSERT INTO supply (coins, height) 
VALUES ($1, $2) 
ON CONFLICT (one_row_id) DO UPDATE 
    SET coins = excluded.coins,
    	height = excluded.height
WHERE supply.height <= excluded.height`

	_, err := db.SQL.Exec(query, pq.Array(dbtypes.NewDbCoins(coins)), height)
	if err != nil {
		return fmt.Errorf("error while storing supply: %s", err)
	}

	return nil
}

// SaveSupply allows to save for the given height the given total amount of coins
func (db *Db) SaveAccountDenomBalance(account string, coin sdk.Coin) error {
	query := `
INSERT INTO account_denom_balance (account, denom, amount) 
VALUES ($1, $2, $3) 
ON CONFLICT (account,denom) DO UPDATE 
    SET amount = $3
WHERE account_denom_balance.account=$1 AND account_denom_balance.denom=$2`

	_, err := db.SQL.Exec(query, account, coin.Denom, coin.Amount.String())
	if err != nil {
		return fmt.Errorf("error while storing account balance: %s", err)
	}

	return nil
}

// SaveSupply allows to save for the given height the given total amount of coins
func (db *Db) DeleteAccountDenomBalance(account string, coin sdk.Coin) error {
	query := `
DELETE FROM account_denom_balance  
WHERE account_denom_balance.account=$1 AND account_denom_balance.denom=$2`

	_, err := db.SQL.Exec(query, account, coin.Denom)
	if err != nil {
		return fmt.Errorf("error while storing account balance: %s", err)
	}

	return nil
}
