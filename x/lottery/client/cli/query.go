package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/naruto0913/lottery/x/lottery/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group lottery queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowLotteryInfo())
	cmd.AddCommand(CmdListStoredLottery())
	cmd.AddCommand(CmdShowStoredLottery())
	cmd.AddCommand(CmdListStoredBet())
	cmd.AddCommand(CmdShowStoredBet())
	cmd.AddCommand(CmdShowBetInfo())
	cmd.AddCommand(CmdShowLotteryData())
	// this line is used by starport scaffolding # 1

	return cmd
}
