package database_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"

	bddbtypes "github.com/forbole/bdjuno/v4/database/types"
)

func (suite *DbTestSuite) TestBigDipperDb_SaveSupply() {
	// Save the data
	original := sdk.NewCoins(
		sdk.NewCoin("desmos", sdk.NewInt(10000)),
		sdk.NewCoin("uatom", sdk.NewInt(15)),
	)
	err := suite.database.SaveSupply(original, 10)
	suite.Require().NoError(err)

	// Verify the data
	expected := bddbtypes.NewSupplyRow(dbtypes.NewDbCoins(original), 10)

	var rows []bddbtypes.SupplyRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM supply`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	coins := sdk.NewCoins(
		sdk.NewCoin("desmos", sdk.NewInt(10000)),
		sdk.NewCoin("uatom", sdk.NewInt(15)),
	)
	err = suite.database.SaveSupply(coins, 9)
	suite.Require().NoError(err)

	// Verify the data
	rows = []bddbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM supply`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with same height
	coins = sdk.NewCoins(sdk.NewCoin("uakash", sdk.NewInt(10)))
	err = suite.database.SaveSupply(coins, 10)
	suite.Require().NoError(err)

	// Verify the data
	expected = bddbtypes.NewSupplyRow(dbtypes.NewDbCoins(coins), 10)

	rows = []bddbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM supply`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with higher height
	coins = sdk.NewCoins(sdk.NewCoin("btc", sdk.NewInt(10)))
	err = suite.database.SaveSupply(coins, 20)
	suite.Require().NoError(err)

	// Verify the data
	expected = bddbtypes.NewSupplyRow(dbtypes.NewDbCoins(coins), 20)

	rows = []bddbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM supply`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_AccoundDenomBalances() {

	// Save the data
	account := "devcore1u6dycnl606n95ggeatusc3zlfd5m4xqpw66et4"
	coin := sdk.NewCoin("ucore", sdk.NewInt(15))
	err := suite.database.SaveAccountDenomBalance(account, coin)
	suite.Require().NoError(err)

	// Verify the data
	expected := bddbtypes.NewAccountBalance(account, coin)

	var rows []bddbtypes.AccountBalance
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM account_denom_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a new value
	coin = sdk.NewCoin("ucore", sdk.NewInt(20))
	err = suite.database.SaveAccountDenomBalance(account, coin)
	suite.Require().NoError(err)

	// Verify the data
	expected = bddbtypes.NewAccountBalance(account, coin)

	rows = []bddbtypes.AccountBalance{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM account_denom_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "supply table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ----------------------------------------------------------------------------------------------------------------

	// Try deleting the value
	err = suite.database.DeleteAccountDenomBalance(account, coin)
	suite.Require().NoError(err)

	// Verify the data
	expected = bddbtypes.NewAccountBalance(account, coin)

	rows = []bddbtypes.AccountBalance{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM account_denom_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0, "supply table should no rows")
}
