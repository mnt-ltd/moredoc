/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/mnt-ltd/daemore"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var (
	user      string
	daemonCmd = &cobra.Command{
		Use:   "daemon",
		Short: "守护进程",
		Long:  `将魔豆文库系统加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	serviceInstallCmd = &cobra.Command{
		Use:   "install",
		Short: "安装魔豆文库服务",
		Long:  `将魔豆文库服务加入到系统服务中，以便在系统启动时自动启动服务。`,
		Run: func(cmd *cobra.Command, args []string) {
			var username []string
			if user != "" {
				username = append(username, user)
			}
			d, err := NewDaemon(username...)
			if err != nil {
				log.Println("魔豆文库服务安装失败：", err)
				return
			}
			err = d.ServiceInstall("serve")
			if err != nil {
				log.Println("魔豆文库服务安装失败：", err)
				return
			}
			log.Println("魔豆文库服务安装成功")
		},
	}

	serviceUninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "卸载魔豆文库服务",
		Long:  `将魔豆文库服务从系统服务中移除，以便在系统启动时不再自动启动服务。`,
		Run: func(cmd *cobra.Command, args []string) {
			d, err := NewDaemon()
			if err != nil {
				log.Println("魔豆文库服务卸载失败：", err)
				return
			}
			err = d.ServiceUninstall()
			if err != nil {
				log.Println("魔豆文库服务卸载失败：", err)
				return
			}
			log.Println("魔豆文库服务卸载成功")
		},
	}

	serviceRestartCmd = &cobra.Command{
		Use:   "restart",
		Short: "重启魔豆文库服务",
		Long:  `重启魔豆文库服务`,
		Run: func(cmd *cobra.Command, args []string) {
			d, err := NewDaemon()
			if err != nil {
				log.Println("魔豆文库服务重启失败：", err)
				return
			}
			err = d.ServiceRestart()
			if err != nil {
				log.Println("魔豆文库服务重启失败：", err)
				return
			}
			log.Println("魔豆文库服务重启成功")
		},
	}

	serviceStartCmd = &cobra.Command{
		Use:   "start",
		Short: "启动魔豆文库服务。",
		Long:  `启动魔豆文库服务。`,
		Run: func(cmd *cobra.Command, args []string) {
			d, err := NewDaemon()
			if err != nil {
				log.Println("魔豆文库服务启动失败：", err)
				return
			}
			err = d.ServiceStart()
			if err != nil {
				log.Println("魔豆文库服务启动失败：", err)
				return
			}
			log.Println("魔豆文库服务启动成功")
		},
	}

	serviceStopCmd = &cobra.Command{
		Use:   "stop",
		Short: "停止魔豆文库服务",
		Long:  `停止魔豆文库服务`,
		Run: func(cmd *cobra.Command, args []string) {
			d, err := NewDaemon()
			if err != nil {
				log.Println("魔豆文库服务停止失败：", err)
				return
			}
			err = d.ServiceStop()
			if err != nil {
				log.Println("魔豆文库服务停止失败：", err)
				return
			}
			log.Println("魔豆文库服务停止成功")
		},
	}

	serviceStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "魔豆文库服务状态",
		Long:  `查看魔豆文库服务的运行状态。`,
		Run: func(cmd *cobra.Command, args []string) {
			d, err := NewDaemon()
			if err != nil {
				log.Println("查看魔豆文库服务状态失败：", err)
				return
			}
			status, err := d.ServiceStatus()
			if err != nil {
				log.Println("查看魔豆文库服务状态失败：", err)
				return
			}
			switch status {
			case service.StatusRunning:
				log.Println("魔豆文库服务状态：正在运行")
			case service.StatusStopped:
				log.Println("魔豆文库服务状态：已停止")
			case service.StatusUnknown:
				log.Println("魔豆文库服务状态：未知")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(daemonCmd)

	serviceInstallCmd.Flags().StringVarP(&user, "user", "u", "", "指定服务运行的用户")
	daemonCmd.AddCommand(serviceInstallCmd)
	daemonCmd.AddCommand(serviceUninstallCmd)
	daemonCmd.AddCommand(serviceRestartCmd)
	daemonCmd.AddCommand(serviceStartCmd)
	daemonCmd.AddCommand(serviceStopCmd)
	daemonCmd.AddCommand(serviceStatusCmd)
}

func NewDaemon(username ...string) (d *daemore.Daemon, err error) {
	d, err = daemore.NewDaemon(daemore.DaemonOption{
		Name:        "moredoc",
		DisplayName: "moredoc service",
		Description: "A document management system, developed using Go language.",
		Callback:    runServer,
	}, username...)
	return
}
