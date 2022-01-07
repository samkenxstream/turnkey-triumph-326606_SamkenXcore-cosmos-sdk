package simulation

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmos/cosmos-sdk/x/group"
)

var initialGroupID = uint64(100000000000000)

// group message types
var (
	TypeMsgCreateGroup                     = sdk.MsgTypeURL(&group.MsgCreateGroup{})
	TypeMsgUpdateGroupMembers              = sdk.MsgTypeURL(&group.MsgUpdateGroupMembers{})
	TypeMsgUpdateGroupAdmin                = sdk.MsgTypeURL(&group.MsgUpdateGroupAdmin{})
	TypeMsgUpdateGroupMetadata             = sdk.MsgTypeURL(&group.MsgUpdateGroupMetadata{})
	TypeMsgCreateGroupPolicy               = sdk.MsgTypeURL(&group.MsgCreateGroupPolicy{})
	TypeMsgUpdateGroupPolicyAdmin          = sdk.MsgTypeURL(&group.MsgUpdateGroupPolicyAdmin{})
	TypeMsgUpdateGroupPolicyDecisionPolicy = sdk.MsgTypeURL(&group.MsgUpdateGroupPolicyDecisionPolicy{})
	TypeMsgUpdateGroupPolicyMetadata       = sdk.MsgTypeURL(&group.MsgUpdateGroupPolicyMetadata{})
	TypeMsgCreateProposal                  = sdk.MsgTypeURL(&group.MsgCreateProposal{})
	TypeMsgVote                            = sdk.MsgTypeURL(&group.MsgVote{})
	TypeMsgExec                            = sdk.MsgTypeURL(&group.MsgExec{})
)

// Simulation operation weights constants
const (
	OpMsgCreateGroup                     = "op_weight_msg_create_group"
	OpMsgUpdateGroupAdmin                = "op_weight_msg_update_group_admin"
	OpMsgUpdateGroupMetadata             = "op_wieght_msg_update_group_metadata"
	OpMsgUpdateGroupMembers              = "op_weight_msg_update_group_members"
	OpMsgCreateGroupPolicy               = "op_weight_msg_create_group_account"
	OpMsgUpdateGroupPolicyAdmin          = "op_weight_msg_update_group_account_admin"
	OpMsgUpdateGroupPolicyDecisionPolicy = "op_weight_msg_update_group_account_decision_policy"
	OpMsgUpdateGroupPolicyMetaData       = "op_weight_msg_update_group_account_metadata"
	OpMsgCreateProposal                  = "op_weight_msg_create_proposal"
	OpMsgVote                            = "op_weight_msg_vote"
	OpMsgExec                            = "ops_weight_msg_exec"
)

// If update group or group account txn's executed, `SimulateMsgVote` & `SimulateMsgExec` txn's returns `noOp`.
// That's why we have less weight for update group & group-account txn's.
const (
	WeightMsgCreateGroup                     = 100
	WeightMsgCreateGroupPolicy               = 100
	WeightMsgCreateProposal                  = 90
	WeightMsgVote                            = 90
	WeightMsgExec                            = 90
	WeightMsgUpdateGroupMetadata             = 5
	WeightMsgUpdateGroupAdmin                = 5
	WeightMsgUpdateGroupMembers              = 5
	WeightMsgUpdateGroupPolicyAdmin          = 5
	WeightMsgUpdateGroupPolicyDecisionPolicy = 5
	WeightMsgUpdateGroupPolicyMetadata       = 5
)

const GroupMemberWeight = 40

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper, appCdc cdctypes.AnyUnpacker) simulation.WeightedOperations {
	var (
		weightMsgCreateGroup                     int
		weightMsgUpdateGroupAdmin                int
		weightMsgUpdateGroupMetadata             int
		weightMsgUpdateGroupMembers              int
		weightMsgCreateGroupPolicy               int
		weightMsgUpdateGroupPolicyAdmin          int
		weightMsgUpdateGroupPolicyDecisionPolicy int
		weightMsgUpdateGroupPolicyMetadata       int
		weightMsgCreateProposal                  int
		weightMsgVote                            int
		weightMsgExec                            int
	)

	appParams.GetOrGenerate(cdc, OpMsgCreateGroup, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = WeightMsgCreateGroup
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateGroupPolicy, &weightMsgCreateGroupPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroupPolicy = WeightMsgCreateGroupPolicy
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = WeightMsgCreateProposal
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = WeightMsgVote
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgExec, &weightMsgExec, nil,
		func(_ *rand.Rand) {
			weightMsgExec = WeightMsgExec
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMetadata, &weightMsgUpdateGroupMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMetadata = WeightMsgUpdateGroupMetadata
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAdmin, &weightMsgUpdateGroupAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAdmin = WeightMsgUpdateGroupAdmin
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMembers, &weightMsgUpdateGroupMembers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMembers = WeightMsgUpdateGroupMembers
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupPolicyAdmin, &weightMsgUpdateGroupPolicyAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupPolicyAdmin = WeightMsgUpdateGroupPolicyAdmin
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupPolicyDecisionPolicy, &weightMsgUpdateGroupPolicyDecisionPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupPolicyDecisionPolicy = WeightMsgUpdateGroupPolicyDecisionPolicy
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupPolicyMetaData, &weightMsgUpdateGroupPolicyMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupPolicyMetadata = WeightMsgUpdateGroupPolicyMetadata
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateGroup,
			SimulateMsgCreateGroup(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateGroupPolicy,
			SimulateMsgCreateGroupPolicy(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProposal,
			SimulateMsgCreateProposal(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgVote,
			SimulateMsgVote(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgExec,
			SimulateMsgExec(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMetadata,
			SimulateMsgUpdateGroupMetadata(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAdmin,
			SimulateMsgUpdateGroupAdmin(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMembers,
			SimulateMsgUpdateGroupMembers(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupPolicyAdmin,
			SimulateMsgUpdateGroupPolicyAdmin(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupPolicyDecisionPolicy,
			SimulateMsgUpdateGroupPolicyDecisionPolicy(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupPolicyMetadata,
			SimulateMsgUpdateGroupPolicyMetadata(ak, bk, k),
		),
	}
}

// SimulateMsgCreateGroup generates a MsgCreateGroup with random values
func SimulateMsgCreateGroup(ak group.AccountKeeper, bk group.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc, _ := simtypes.RandomAcc(r, accounts)
		account := ak.GetAccount(ctx, acc.Address)
		accAddr := acc.Address.String()

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, "fee error"), nil, err
		}

		members := []group.Member{
			{
				Address:  accAddr,
				Weight:   fmt.Sprintf("%d", GroupMemberWeight),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := &group.MsgCreateGroup{Admin: accAddr, Members: members, Metadata: []byte(simtypes.RandStringOfLength(r, 10))}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroup, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

// SimulateMsgCreateGroupPolicy generates a NewMsgCreateGroupPolicy with random values
func SimulateMsgCreateGroupPolicy(ak group.AccountKeeper, bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		groupID, acc, account, err := randomGroup(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, ""), nil, err
		}
		if groupID == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, ""), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, "fee error"), nil, err
		}

		msg, err := group.NewMsgCreateGroupPolicy(
			acc.Address,
			groupID,
			[]byte(simtypes.RandStringOfLength(r, 10)),
			&group.ThresholdDecisionPolicy{
				Threshold: "20",
				Timeout:   time.Second * time.Duration(30*24*60*60),
			},
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			fmt.Printf("ERR DELIVER %v\n", err)
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

// SimulateMsgCreateProposal generates a NewMsgCreateProposal with random values
func SimulateMsgCreateProposal(ak group.AccountKeeper, bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		groupID, groupPolicyAddr, _, _, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no group policy found"), nil, nil
		}

		// Pick a random member from the group
		res, err := k.GroupMembers(sdk.WrapSDKContext(sdkCtx), &group.QueryGroupMembersRequest{
			GroupId: groupID,
		})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, ""), nil, err
		}
		n := randIntInRange(r, len(res.Members))
		if n < 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no group member found"), nil, nil
		}
		idx := findAccount(accounts, res.Members[n].Member.Address)
		if idx < 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no account found for group member"), nil, nil
		}
		acc := accounts[idx]
		account := ak.GetAccount(sdkCtx, acc.Address)

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "fee error"), nil, err
		}

		msg := group.MsgCreateProposal{
			Address:   groupPolicyAddr,
			Proposers: []string{acc.Address.String()},
			Metadata:  []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		// err = msg.UnpackInterfaces(appCdc)
		// if err != nil {
		// 	return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "unmarshal error"), nil, err
		// }
		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// SimulateMsgUpdateGroupAdmin generates a MsgUpdateGroupAdmin with random values
func SimulateMsgUpdateGroupAdmin(ak group.AccountKeeper, bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		groupID, acc, account, err := randomGroup(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, ""), nil, err
		}
		if groupID == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, "no group found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, "fee error"), nil, err
		}

		if len(accounts) == 1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, "can't set a new admin with only one account"), nil, nil
		}
		newAdmin, _ := simtypes.RandomAcc(r, accounts)
		// disallow setting current admin as new admin
		for acc.PubKey.Equals(newAdmin.PubKey) {
			newAdmin, _ = simtypes.RandomAcc(r, accounts)
		}

		msg := group.MsgUpdateGroupAdmin{
			GroupId:  groupID,
			Admin:    account.GetAddress().String(),
			NewAdmin: newAdmin.Address.String(),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// SimulateMsgUpdateGroupMetadata generates a MsgUpdateGroupMetadata with random values
func SimulateMsgUpdateGroupMetadata(ak group.AccountKeeper, bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		groupID, acc, account, err := randomGroup(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, ""), nil, err
		}
		if groupID == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, "no group found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMetadata, "fee error"), nil, err
		}

		msg := group.MsgUpdateGroupMetadata{
			GroupId:  groupID,
			Admin:    account.GetAddress().String(),
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// SimulateMsgUpdateGroupMembers generates a MsgUpdateGroupMembers with random values
func SimulateMsgUpdateGroupMembers(ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		groupID, acc, account, err := randomGroup(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, ""), nil, err
		}
		if groupID == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateGroupPolicy, "no group found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMembers, "fee error"), nil, err
		}

		member, _ := simtypes.RandomAcc(r, accounts)

		members := []group.Member{
			{
				Address:  member.Address.String(),
				Weight:   fmt.Sprintf("%d", GroupMemberWeight),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := group.MsgUpdateGroupMembers{
			GroupId:       groupID,
			Admin:         acc.Address.String(),
			MemberUpdates: members,
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupMembers, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// SimulateMsgUpdateGroupPolicyAdmin generates a MsgUpdateGroupPolicyAdmin with random values
func SimulateMsgUpdateGroupPolicyAdmin(ak group.AccountKeeper, bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, groupPolicyAddr, acc, account, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no group policy found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyAdmin, "fee error"), nil, err
		}

		if len(accounts) == 1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyAdmin, "can't set a new admin with only one account"), nil, nil
		}
		newAdmin, _ := simtypes.RandomAcc(r, accounts)
		// disallow setting current admin as new admin
		for acc.PubKey.Equals(newAdmin.PubKey) {
			newAdmin, _ = simtypes.RandomAcc(r, accounts)
		}

		msg := group.MsgUpdateGroupPolicyAdmin{
			Admin:    acc.Address.String(),
			Address:  groupPolicyAddr,
			NewAdmin: newAdmin.Address.String(),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// // SimulateMsgUpdateGroupPolicyDecisionPolicy generates a NewMsgUpdateGroupPolicyDecisionPolicyRequest with random values
func SimulateMsgUpdateGroupPolicyDecisionPolicy(ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, groupPolicyAddr, acc, account, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no group policy found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyDecisionPolicy, "fee error"), nil, err
		}

		groupPolicyBech32, err := sdk.AccAddressFromBech32(groupPolicyAddr)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		msg, err := group.NewMsgUpdateGroupPolicyDecisionPolicyRequest(acc.Address, groupPolicyBech32, &group.ThresholdDecisionPolicy{
			Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 20)),
			Timeout:   time.Second * time.Duration(simtypes.RandIntBetween(r, 100, 1000)),
		})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyDecisionPolicy, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyDecisionPolicy, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

// // SimulateMsgUpdateGroupPolicyMetadata generates a MsgUpdateGroupPolicyMetadata with random values
func SimulateMsgUpdateGroupPolicyMetadata(ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, groupPolicyAddr, acc, account, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgCreateProposal, "no group policy found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyMetadata, "fee error"), nil, err
		}

		msg := group.MsgUpdateGroupPolicyMetadata{
			Admin:    acc.Address.String(),
			Address:  groupPolicyAddr,
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// SimulateMsgVote generates a MsgVote with random values
func SimulateMsgVote(ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, groupPolicyAddr, acc1, account, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no group policy found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fee error"), nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)
		proposalsResult, err := k.ProposalsByGroupPolicy(ctx, &group.QueryProposalsByGroupPolicyRequest{Address: groupPolicyAddr})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fail to query group info"), nil, err
		}
		proposals := proposalsResult.GetProposals()
		if len(proposals) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		proposalID := -1

		for _, proposal := range proposals {
			if proposal.Status == group.ProposalStatusSubmitted {
				timeout := proposal.Timeout
				if err != nil {
					return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "error: while converting to timestamp"), nil, err
				}
				proposalID = int(proposal.ProposalId)
				if timeout.Before(sdkCtx.BlockTime()) || timeout.Equal(sdkCtx.BlockTime()) {
					return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "voting period ended: skipping"), nil, nil
				}
				break
			}
		}

		// return no-op if no proposal found
		if proposalID == -1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		msg := group.MsgVote{
			ProposalId: uint64(proposalID),
			Voter:      acc1.Address.String(),
			Choice:     group.Choice_CHOICE_YES,
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)

		if err != nil {
			if strings.Contains(err.Error(), "group was modified") || strings.Contains(err.Error(), "group account was modified") {
				return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "no-op:group/group-account was modified"), nil, nil
			}
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

// // SimulateMsgExec generates a MsgExec with random values
func SimulateMsgExec(ak group.AccountKeeper,
	bk group.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, groupPolicyAddr, acc1, account, err := randomGroupPolicy(r, k, ak, sdkCtx, accounts)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, ""), nil, err
		}
		if groupPolicyAddr == "" {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no group policy found"), nil, nil
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fee error"), nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)
		proposalsResult, err := k.ProposalsByGroupPolicy(ctx, &group.QueryProposalsByGroupPolicyRequest{Address: groupPolicyAddr})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "fail to query group info"), nil, err
		}
		proposals := proposalsResult.GetProposals()
		if len(proposals) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		proposalID := -1

		for _, proposal := range proposals {
			if proposal.Status == group.ProposalStatusClosed {
				proposalID = int(proposal.ProposalId)
				break
			}
		}

		// return no-op if no proposal found
		if proposalID == -1 {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgVote, "no proposals found"), nil, nil
		}

		msg := group.MsgExec{
			ProposalId: uint64(proposalID),
			Signer:     acc1.Address.String(),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, TypeMsgUpdateGroupPolicyMetadata, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "group was modified") || strings.Contains(err.Error(), "group account was modified") {
				return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "no-op:group/group-account was modified"), nil, nil
			}
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(&msg, true, "", nil), nil, err
	}
}

func randomGroup(r *rand.Rand, k keeper.Keeper, ak group.AccountKeeper,
	ctx sdk.Context, accounts []simtypes.Account) (groupID uint64, acc simtypes.Account, account authtypes.AccountI, err error) {
	groupID = k.GetGroupSequence(ctx)

	switch {
	case groupID > initialGroupID:
		// select a random ID between [initialGroupID, groupID]
		groupID = uint64(simtypes.RandIntBetween(r, int(initialGroupID), int(groupID)))

	default:
		// This is called on the first call to this function
		// in order to update the global variable
		initialGroupID = groupID
	}

	res, err := k.GroupInfo(sdk.WrapSDKContext(ctx), &group.QueryGroupInfoRequest{GroupId: groupID})
	if err != nil {
		return 0, simtypes.Account{}, nil, err
	}

	groupAdmin := res.Info.Admin
	found := -1
	for i := range accounts {
		if accounts[i].Address.String() == groupAdmin {
			found = i
			break
		}
	}
	if found < 0 {
		return 0, simtypes.Account{}, nil, nil
	}
	acc = accounts[found]
	account = ak.GetAccount(ctx, acc.Address)
	return groupID, acc, account, nil
}

func randomGroupPolicy(r *rand.Rand, k keeper.Keeper, ak group.AccountKeeper,
	ctx sdk.Context, accounts []simtypes.Account) (groupID uint64, groupPolicyAddr string, acc simtypes.Account, account authtypes.AccountI, err error) {
	groupID, _, _, err = randomGroup(r, k, ak, ctx, accounts)
	if err != nil {
		return 0, "", simtypes.Account{}, nil, err
	}
	if groupID == 0 {
		return 0, "", simtypes.Account{}, nil, nil
	}

	result, err := k.GroupPoliciesByGroup(sdk.WrapSDKContext(ctx), &group.QueryGroupPoliciesByGroupRequest{GroupId: groupID})
	if err != nil {
		return groupID, "", simtypes.Account{}, nil, err
	}

	n := randIntInRange(r, len(result.GroupPolicies))
	if n < 0 {
		return groupID, "", simtypes.Account{}, nil, nil
	}
	groupPolicy := result.GroupPolicies[n]
	groupPolicyAddr = groupPolicy.Address

	idx := findAccount(accounts, groupPolicy.Admin)
	if idx < 0 {
		return groupID, "", simtypes.Account{}, nil, nil
	}
	acc = accounts[idx]
	account = ak.GetAccount(ctx, acc.Address)
	return groupID, groupPolicyAddr, acc, account, nil
}

func randIntInRange(r *rand.Rand, l int) int {
	if l == 0 {
		return -1
	}
	if l == 1 {
		return 0
	} else {
		return simtypes.RandIntBetween(r, 0, l-1)
	}
}

func findAccount(accounts []simtypes.Account, addr string) (idx int) {
	idx = -1
	for i := range accounts {
		if accounts[i].Address.String() == addr {
			idx = i
			break
		}
	}
	return idx
}
