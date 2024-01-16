/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"moredoc/service"

	"github.com/spf13/cobra"
)

// migrateContentCmd represents the migrateContent command
var migrateContentCmd = &cobra.Command{
	Use:   "migrateContent",
	Short: "迁移内容",
	Long:  `将文档文本内容迁移到数据库中`,
	Run: func(cmd *cobra.Command, args []string) {
		service.MigrateContent(cfg, logger)
	},
}

func init() {
	rootCmd.AddCommand(migrateContentCmd)
}
