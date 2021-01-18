package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cdptypes "github.com/kava-labs/kava/x/cdp/types"
	hardtypes "github.com/kava-labs/kava/x/hard/types"
)

// Hooks wrapper struct for hooks
type Hooks struct {
	k Keeper
}

var _ cdptypes.CDPHooks = Hooks{}
var _ hardtypes.HARDHooks = Hooks{}

// Hooks create new incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterCDPCreated function that runs after a cdp is created
func (h Hooks) AfterCDPCreated(ctx sdk.Context, cdp cdptypes.CDP) {
	h.k.InitializeUSDXMintingClaim(ctx, cdp)
}

// BeforeCDPModified function that runs before a cdp is modified
// note that this is called immediately after interest is synchronized, and so could potentially
// be called AfterCDPInterestUpdated or something like that, if we we're to expand the scope of cdp hooks
func (h Hooks) BeforeCDPModified(ctx sdk.Context, cdp cdptypes.CDP) {
	h.k.SynchronizeUSDXMintingReward(ctx, cdp)
}

// BeforeDepositCreated function that runs before a deposit is created
func (h Hooks) BeforeDepositCreated(ctx sdk.Context, deposit hardtypes.Deposit, denom string) {
	h.k.InitializeHardSupplyReward(ctx, deposit, denom)
}

// BeforeDepositModified function that runs before a deposit is modified
func (h Hooks) BeforeDepositModified(ctx sdk.Context, deposit hardtypes.Deposit, denom string) {
	h.k.SynchronizeHardSupplyReward(ctx, deposit, denom)
}

// AfterDepositModified function that runs after a deposit is modified
func (h Hooks) AfterDepositModified(ctx sdk.Context, deposit hardtypes.Deposit) {
	// TODO: Add any new denoms in the Deposit to SupplyIndexes
}

// BeforeBorrowCreated function that runs before a borrow is created
func (h Hooks) BeforeBorrowCreated(ctx sdk.Context, borrow hardtypes.Borrow, denom string) {
	h.k.InitializeHardBorrowReward(ctx, borrow, denom)
}

// BeforeBorrowModified function that runs before a borrow is modified
func (h Hooks) BeforeBorrowModified(ctx sdk.Context, borrow hardtypes.Borrow, denom string) {
	h.k.SynchronizeHardBorrowReward(ctx, borrow, denom)
}

// AfterBorrowModified function that runs after a borrow is modified
func (h Hooks) AfterBorrowModified(ctx sdk.Context, borrow hardtypes.Borrow) {
	// TODO: Add any new denoms in the Borrow to BorrowIndexes
}
