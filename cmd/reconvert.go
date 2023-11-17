/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"moredoc/service"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ext        string // 指定的文件后缀，如：png、jpg、webp
	documentId int64  // 指定的文档ID, 0表示全部，大于0表示指定文档
)

// reconvertCmd represents the reconvert command
var reconvertCmd = &cobra.Command{
	Use:   "reconvert",
	Short: "文档重转",
	Long: `将已转换成功的文档重新转为png、jpg或者webp格式，以便于提高预览速度。
【注意】该指令只对已转换成功的文档有效，且只能转为png、jpg或者webp格式中的一种。
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// ext必须是png、jpg、webp中的一种
		ext = strings.ToLower(ext)
		if ext != "png" && ext != "jpg" && ext != "webp" {
			ext = "webp"
		}

		// 必须指定documentId
		if documentId < 0 {
			fmt.Println("\n请用--id指定的文档ID, 0表示全部，大于0表示指定文档。如需重转全部文档，建议先指定一个文档进行测试查验效果是否符合需求。")
			fmt.Println("\n按回车键退出...")
			fmt.Scanln()
			return
		}

		// 提示用户，输入Y确认，否则退出
		fmt.Println("\n请确认是否重转文档？")
		if documentId == 0 {
			fmt.Print("转换文档：全部文档")
		} else {
			fmt.Print("转换文档ID：", documentId)
		}
		fmt.Println("；文档预览格式：", ext)
		fmt.Println("按 Y 确认，按其他键取消和退出...")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "Y" && confirm != "y" {
			fmt.Println("\n已取消重转文档。")
			return
		}

		service.Reconvert(cfg, logger, ext, documentId)
	},
}

func init() {
	rootCmd.AddCommand(reconvertCmd)

	reconvertCmd.Flags().StringVarP(&ext, "ext", "e", "webp", "指定的文档预览格式，如：png、jpg、webp")
	reconvertCmd.Flags().Int64VarP(&documentId, "id", "d", -1, "指定的文档ID, 0表示全部，大于0表示指定文档。如需重转全部文档，建议先指定一个文档进行测试查验效果是否符合需求。")
}
