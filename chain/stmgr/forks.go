package stmgr

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"go.opencensus.io/trace"
)

var ForksAtHeight = map[abi.ChainEpoch]func(context.Context, *StateManager, types.StateTree) error{}

func (sm *StateManager) handleStateForks(ctx context.Context, st types.StateTree, height abi.ChainEpoch) (err error) {
	ctx, span := trace.StartSpan(ctx, "stmgr.handleStateForks")
	defer span.End()

	f, ok := ForksAtHeight[height]
	if ok {
		err := f(ctx, sm, st)
		if err != nil {
			return err
		}
	}

	return nil
}
