package keeper

import (
	"etherlink/x/etherlink/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEthState set ethState in the store
func (k Keeper) SetEthState(ctx sdk.Context, ethState types.EthState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKey))
	b := k.cdc.MustMarshal(&ethState)
	store.Set([]byte{0}, b)
}

// GetEthState returns ethState
func (k Keeper) GetEthState(ctx sdk.Context) (val types.EthState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthState removes ethState from the store
func (k Keeper) RemoveEthState(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKey))
	store.Delete([]byte{0})
}
