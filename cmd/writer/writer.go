/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package writer

import (
	"github.com/spf13/cobra"
)

// writerCmd represents the writer command
var writerCmd = &cobra.Command{
	Use:   "writer",
	Short: "Writer operations",
	Long:  `Writer operations`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var (
	name string
)

func InitCmds(rootCmd *cobra.Command) {
	rootCmd.AddCommand(writerCmd)
}
