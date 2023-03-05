/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"moredoc/service"
	"moredoc/util"
	"moredoc/util/command"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a server",
	Long:  `start a server`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Version = Version
		util.Hash = GitHash
		util.BuildAt = BuildAt

		c := make(chan os.Signal, 1)
		// 监听退出信号
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			//阻塞直至有信号传入
			s := <-c
			// 收到退出信号，关闭子进程
			fmt.Println("get signal：", s)
			fmt.Println("close child process...")
			command.CloseChildProccess()
			fmt.Println("close child process done.")
			fmt.Println("exit.")
			os.Exit(0)
		}()
		service.Run(cfg, logger)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
