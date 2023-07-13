package keeper

import (
	"bytes"
	"context"
	"encoding/hex"

	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/or/settlement/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (s msgServer) StartChallenge(c context.Context, req *types.MsgStartChallenge) (*types.MsgStartChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	startState, _ := hex.DecodeString("1dd6bce3200aab85eefdcf1b20f2f61601f14f3804ed984f243bd42eef6d387c") // start state of miniapp (demo application)
	// TODO:
	// - get start state from rollup module
	// - validate if rollup name exist
	// - validate if challenger exist in rollup
	// - validate if defender exist in rollup
	// - validate if block height exist in rollup blocks
	challenge := types.Challenge{
		L:                   0,
		R:                   req.StepCount,
		AssertedStateHashes: map[uint64][]byte{0: startState[:]},
		DefendedStateHashes: map[uint64][]byte{0: startState[:]},
		Challenger:          req.From,
		Defender:            req.To,
		BlockHeight:         req.BlockHeight,
		RollupName:          req.RollupName,
	}

	if s.Keeper.HasChallenge(ctx, challenge.ID()) {
		return nil, types.ErrChallengeExists
	}
	s.Keeper.SetChallenge(ctx, challenge.ID(), challenge)

	return &types.MsgStartChallengeResponse{
		ChallengeId: challenge.ID(),
	}, nil
}

func (s msgServer) NsectChallenge(c context.Context, req *types.MsgNsectChallenge) (*types.MsgNsectChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	challenge, err := s.Keeper.GetChallenge(ctx, req.ChallengeId)
	if err != nil {
		return nil, err
	}

	if !challenge.IsSearching() {
		return nil, types.ErrNotSearching
	}

	if challenge.CurrentResponder() != req.From {
		return nil, types.ErrNotResponder
	}

	if len(challenge.GetSteps()) != len(req.StateHashes) {
		return nil, types.ErrInvalidStateHashes
	}

	steps := challenge.GetSteps()

	if challenge.Challenger == req.From {
		for i := range steps {
			challenge.AssertedStateHashes[steps[i]] = req.StateHashes[i]
		}
	} else {
		for i := range steps {
			challenge.DefendedStateHashes[steps[i]] = req.StateHashes[i]
		}

		// next round
		for i := range steps {
			if bytes.Equal(challenge.AssertedStateHashes[steps[i]], challenge.DefendedStateHashes[steps[i]]) {
				challenge.L = steps[i]
			} else {
				challenge.R = steps[i]
				break
			}
		}
	}

	s.Keeper.SetChallenge(ctx, challenge.ID(), *challenge)

	return &types.MsgNsectChallengeResponse{}, nil
}

func (s msgServer) FinishChallenge(c context.Context, req *types.MsgFinishChallenge) (*types.MsgFinishChallengeResponse, error) {
	panic("implement me")
}
