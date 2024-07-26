package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
var _ DistributionHook = &MultiDistributionHooks{}

type MultiDistributionHooks []DistributionHook

func NewMultiDistributionHooks(hooks ...DistributionHook) MultiDistributionHooks {
	return hooks
}

func (h MultiDistributionHooks) BeforeFeeCollectorSend(ctx context.Context, feeCollector sdk.ModuleAccountI) error {
	for i := range h {
		if err := h[i].BeforeFeeCollectorSend(ctx, feeCollector); err != nil {
			return err
		}
	}

	return nil
}
