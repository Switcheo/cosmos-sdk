package types

import (
	context "context"

	"cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the account contract that must be fulfilled when
// creating a x/bank keeper.
type AccountKeeper interface {
	AddressCodec() address.Codec

	NewAccount(context.Context, sdk.AccountI) sdk.AccountI
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI

	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetAllAccounts(ctx context.Context) []sdk.AccountI
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	SetAccount(ctx context.Context, acc sdk.AccountI)

	IterateAccounts(ctx context.Context, process func(sdk.AccountI) bool)

	ValidatePermissions(macc sdk.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx context.Context, moduleName string) (sdk.ModuleAccountI, []string)
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
	SetModuleAccount(ctx context.Context, macc sdk.ModuleAccountI)
	GetModulePermissions() map[string]types.PermissionsForAddress

	GetMergedAccountAddressIfExists(ctx context.Context, addr sdk.AccAddress) sdk.AccAddress
	GetMappedAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccAddress
}
type SendHooks interface {
	BeforeSend(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	AfterSend(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	BeforeMultiSend(ctx context.Context, inputs []Input, outputs []Output) error
	AfterMultiSend(ctx context.Context, inputs []Input, outputs []Output) error
}
