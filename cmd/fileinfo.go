/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zero-gravity-labs/zerog-storage-tool/core"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "fileinfo",
	Short: "Get file info by hash",
	Long:  `Get file info by hash`,
	Run: func(cmd *cobra.Command, args []string) {
		fi, err := core.GetFileInfo(common.HexToHash(root))
		if err != nil {
			logrus.WithError(err).Error("Failed to get file info")
		} else {
			logrus.WithField("fi", fi).Info("Get file info completed")
		}
	},
}

var (
	root string
)

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVarP(&root, "root", "r", "", "root hash of content")
	queryCmd.MarkFlagRequired("root")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
