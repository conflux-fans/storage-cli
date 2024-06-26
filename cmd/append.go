/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/conflux-fans/storage-cli/core"
	"github.com/conflux-fans/storage-cli/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append content to specified file",
	Long:  `Append content to specified file`,
	Run: func(cmd *cobra.Command, args []string) {
		if !common.IsHexAddress(account) {
			logger.Failf("account %s is not valid address", account)
			return
		}

		if data != "" {
			if err := core.AppendData(common.HexToAddress(account), name, data); err != nil {
				logger.Fail(err.Error())
			}
			return
		}

		if filePath != "" {
			if err := core.AppendFromFile(common.HexToAddress(account), name, filePath); err != nil {
				logger.Fail(err.Error())
			}
			return
		}
	},
}

var (
	data    string
	account string
)

func init() {
	rootCmd.AddCommand(appendCmd)
	appendCmd.Flags().StringVar(&filePath, "file", "", "file path of content to upload")
	appendCmd.Flags().StringVar(&data, "data", "", "append content")
	appendCmd.Flags().StringVar(&name, "name", "", "name, for appending content")
	appendCmd.Flags().StringVar(&account, "account", "", "name, for appending content")
	appendCmd.MarkFlagsOneRequired("data", "file")
}
