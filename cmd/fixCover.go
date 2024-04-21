/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"moredoc/service"

	"github.com/spf13/cobra"
)

// fixCoverCmd represents the fixCover command
var fixCoverCmd = &cobra.Command{
	Use:   "fixCover",
	Short: "修正封面大小",
	Long:  `修正已有图片的封面大小，使其符合要求，特别是PPT类的文档。`,
	Run: func(cmd *cobra.Command, args []string) {
		service.FixCover(cfg, logger)
	},
}

func init() {
	rootCmd.AddCommand(fixCoverCmd)
}
