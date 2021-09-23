package middleware_test

import (
	"github.com/tendermint/tendermint/abci/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/middleware"
)

func (suite *MWTestSuite) TestRunMsgs() {
	ctx := suite.SetupTest(true) // setup

	msr := middleware.NewMsgServiceRouter(suite.clientCtx.InterfaceRegistry)
	testdata.RegisterMsgServer(msr, testdata.MsgServerImpl{})
	txHandler := middleware.NewRunMsgsTxHandler(msr, nil)

	priv, _, _ := testdata.KeyTestPubAddr()
	txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(&testdata.MsgCreateDog{Dog: &testdata.Dog{Name: "Spot"}})
	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv}, []uint64{0}, []uint64{0}
	tx, _, err := suite.createTestTx(txBuilder, privs, accNums, accSeqs, ctx.ChainID())
	suite.Require().NoError(err)
	txBytes, err := suite.clientCtx.TxConfig.TxEncoder()(tx)
	suite.Require().NoError(err)

	res, err := txHandler.DeliverTx(sdk.WrapSDKContext(ctx), tx, types.RequestDeliverTx{Tx: txBytes})
	suite.Require().NoError(err)
	suite.Require().NotEmpty(res.Data)
	var txMsgData sdk.TxMsgData
	err = suite.clientCtx.Codec.Unmarshal(res.Data, &txMsgData)
	suite.Require().NoError(err)
	suite.Require().Len(txMsgData.Data, 1)
	suite.Require().Equal(sdk.MsgTypeURL(&testdata.MsgCreateDog{}), txMsgData.Data[0].MsgType)
}
