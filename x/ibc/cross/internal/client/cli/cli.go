package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "cross",
		Short: "Cross chain contract calls transaction subcommands",
	}

	txCmd.AddCommand(flags.PostCommands(
		GetInitiateTxCmd(cdc),
		GetCreateContractTransactionFileCmd(cdc),
		GetMergeContractTransactionFilesCmd(cdc),
	)...)

	return txCmd
}