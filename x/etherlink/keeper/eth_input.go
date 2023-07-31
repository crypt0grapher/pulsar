package keeper

import (
	"etherlink/x/etherlink/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEthInput set ethInput in the store
func (k Keeper) SetEthInput(ctx sdk.Context, ethInput types.EthInput) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthInputKey))
	b := k.cdc.MustMarshal(&ethInput)
	store.Set([]byte{0}, b)
}

// GetEthInput returns ethInput
func (k Keeper) GetEthInput(ctx sdk.Context) (val types.EthInput, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthInputKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthInput removes ethInput from the store
func (k Keeper) RemoveEthInput(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthInputKey))
	store.Delete([]byte{0})
}
