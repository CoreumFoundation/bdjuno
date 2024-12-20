package utils

import (
	"sync"

	"cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	"github.com/cosmos/ibc-go/modules/capability"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	assetft "github.com/CoreumFoundation/coreum/v5/x/asset/ft"
	assetnft "github.com/CoreumFoundation/coreum/v5/x/asset/nft"
	"github.com/CoreumFoundation/coreum/v5/x/customparams"
	"github.com/CoreumFoundation/coreum/v5/x/delay"
	"github.com/CoreumFoundation/coreum/v5/x/dex"
	"github.com/CoreumFoundation/coreum/v5/x/feemodel"
	"github.com/CoreumFoundation/coreum/v5/x/wbank"
	"github.com/CoreumFoundation/coreum/v5/x/wibctransfer"
	"github.com/CoreumFoundation/coreum/v5/x/wnft"
)

var (
	once sync.Once
	cdc  *codec.ProtoCodec
)

func GetCodec() codec.Codec {
	once.Do(func() {
		interfaceRegistry := codectypes.NewInterfaceRegistry()
		getBasicManagers().RegisterInterfaces(interfaceRegistry)
		std.RegisterInterfaces(interfaceRegistry)
		cdc = codec.NewProtoCodec(interfaceRegistry)
	})
	return cdc
}

// getBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func getBasicManagers() module.BasicManager {
	return module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		wbank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		ibchooks.AppModuleBasic{},
		packetforward.AppModuleBasic{},
		ica.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		wibctransfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		wasm.AppModuleBasic{},
		feemodel.AppModuleBasic{},
		wnft.AppModuleBasic{},
		assetft.AppModuleBasic{},
		assetnft.AppModuleBasic{},
		customparams.AppModuleBasic{},
		delay.AppModuleBasic{},
		dex.AppModuleBasic{},
	)
}

// UnpackMessage unpacks a message from a byte slice
func UnpackMessage[T proto.Message](cdc codec.Codec, bz []byte, ptr T) T {
	var any codectypes.Any
	cdc.MustUnmarshalJSON(bz, &any)
	var cosmosMsg sdk.Msg
	if err := cdc.UnpackAny(&any, &cosmosMsg); err != nil {
		panic(err)
	}
	return cosmosMsg.(T)
}
