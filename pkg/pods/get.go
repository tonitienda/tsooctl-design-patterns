package pods

import "tsoobame/tsooctl/pkg/contracts"

type GetPods struct {
}

func (getPods GetPods) Execute() contracts.CommandResult {
	return contracts.CommandResult{
		Message: "Getting pods.",
	}
}
