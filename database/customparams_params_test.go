package database_test

import (
	"encoding/json"

	customparamstypes "github.com/CoreumFoundation/coreum/v2/x/customparams/types"
	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/forbole/bdjuno/v3/types"
)

func (suite *DbTestSuite) TestSaveCustomParamsParams() {
	customParamsDefaultStakingParams := customparamstypes.DefaultStakingParams()
	err := suite.database.SaveCustomParamsParams(
		types.NewCustomParamsParams(types.CustomParamsStakingParams(customParamsDefaultStakingParams), 10),
	)
	suite.Require().NoError(err)

	var rows []dbtypes.CustomParamsParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM customparams_params`)
	suite.Require().NoError(err)

	suite.Require().Len(rows, 1)

	var storedStakingParams customparamstypes.StakingParams
	err = json.Unmarshal([]byte(rows[0].StakingParams), &storedStakingParams)
	suite.Require().NoError(err)
	suite.Require().Equal(customParamsDefaultStakingParams, storedStakingParams)
}

func (suite *DbTestSuite) TestGetCustomParamsParams() {
	customparamsParams := customparamstypes.DefaultStakingParams()

	paramsBz, err := json.Marshal(&customparamsParams)
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(
		`INSERT INTO customparams_params (staking_params, height) VALUES ($1, $2)`,
		string(paramsBz), 10,
	)
	suite.Require().NoError(err)

	params, err := suite.database.GetCustomParamsParams()
	suite.Require().NoError(err)

	suite.Require().Equal(&types.CustomParamsParams{
		StakingParams: types.CustomParamsStakingParams(customparamsParams),
		Height:        10,
	}, params)
}
