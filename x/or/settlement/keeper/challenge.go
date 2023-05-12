package keeper

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/or/settlement/types"
)

func (k Keeper) SetChallenge(ctx sdk.Context, challengeID int64, challenge types.Challenge) {
	store := ctx.KVStore(k.storeKey)
	key := challengeKey(challengeID)

	bz, err := challenge.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
}

func (k Keeper) GetChallenge(ctx sdk.Context, challengeID int64) (*types.Challenge, error) {
	store := ctx.KVStore(k.storeKey)
	key := challengeKey(challengeID)
	bz := store.Get(key)
	if bz == nil {
		return nil, types.ErrChallengeNotExist.Wrapf("no challenge for %d", challengeID)
	}

	var challenge types.Challenge
	k.cdc.MustUnmarshal(bz, &challenge)

	return &challenge, nil
}
