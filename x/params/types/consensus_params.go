package types

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
)

// ConsensusParamsKeyTable returns an x/params module keyTable to be used in
// the BaseApp's ParamStore. The KeyTable registers the types along with the
// standard validation functions. Applications can choose to adopt this KeyTable
// or provider their own when the existing validation functions do not suite their
// needs.
func ConsensusParamsKeyTable() KeyTable {
	return NewKeyTable(
		NewParamSetPair(
			baseapp.ParamStoreKeyBlockParams, tmproto.BlockParams{}, baseapp.ValidateBlockParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyEvidenceParams, tmproto.EvidenceParams{}, baseapp.ValidateEvidenceParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyValidatorParams, tmproto.ValidatorParams{}, baseapp.ValidateValidatorParams,
		),
		// This param is stroed in the param store in order to send it to tendermint after state sync. If this param is modified by governance, it will be reset by the upgrade module
		NewParamSetPair(
			baseapp.ParamStoreKeyVersionParams, tmproto.VersionParams{}, baseapp.ValidateValidatorParams,
		),
	)
}
