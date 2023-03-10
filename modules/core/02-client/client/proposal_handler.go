package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/vipernet-xyz/ibc-go/v7/modules/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
)
