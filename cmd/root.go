/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"moredoc/conf"
	"moredoc/service"
	"moredoc/util"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mnt-ltd/command"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	cfgFile   string
	cfg       = &conf.Config{}
	logger, _ = zap.NewProduction()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "moredoc",
	Short: "魔豆文库，文库系统解决方案",
	Long:  `魔豆文库，使用Go语言开发的类似百度文库、新浪爱问文库的文库系统解决方案，支持 TXT、PDF、EPUB、MOBI、Office 等格式文档的在线预览与管理，为 dochub文库的重构版本。`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	os.Chdir(filepath.Dir(os.Args[0]))
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/app.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "app" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(filepath.Dir(os.Args[0]))
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
	}

	viper.SetEnvPrefix("MOREDOC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Println("viper.Unmarshal", err, viper.ConfigFileUsed())
	}

	overwriteConfig(cfg)

	initLogger(cfg.Level, cfg.LogEncoding, cfg.Logger)

	cfg.Database.Prefix = "mnt_"

	logger.Info("config", zap.String("Using config file:", viper.ConfigFileUsed()), zap.Any("config", cfg))
}

func getEnvDefaultString(key string, df1 string, df2 ...string) string {
	val := viper.GetString(key)
	if val == "" {
		val = df1
	}
	if val == "" && len(df2) > 0 {
		val = df2[0]
	}
	return val
}

func getEnvDefaultBool(key string, df1 bool) bool {
	val := viper.GetString(key)
	if val != "" {
		return strings.ToUpper(val) == "TRUE"
	}
	return df1
}

func getEnvDefaultInt64(key string, df1 int64, df2 ...int64) int64 {
	v := viper.GetInt64(key)
	if v > 0 {
		return v
	}
	if df1 > 0 {
		return df1
	}
	if len(df2) > 0 {
		return df2[0]
	}
	return 0
}

func overwriteConfig(cfg *conf.Config) {
	// 基础配置
	cfg.Level = getEnvDefaultString("LEVEL", cfg.Level, "info")
	cfg.LogEncoding = getEnvDefaultString("LOG_ENCODING", cfg.LogEncoding, "console")
	cfg.Port = int(getEnvDefaultInt64("PORT", int64(cfg.Port), 8880))

	// JWT
	cfg.JWT.Secret = getEnvDefaultString("JWT_SECRET", cfg.JWT.Secret)
	cfg.JWT.ExpireDays = getEnvDefaultInt64("JWT_EXPIRE_DAYS", cfg.JWT.ExpireDays, 365)

	// database
	cfg.Database.DSN = getEnvDefaultString("DATABASE_DSN", cfg.Database.DSN)
	cfg.Database.ShowSQL = getEnvDefaultBool("DATABASE_SHOW_SQL", cfg.Database.ShowSQL)
	cfg.Database.MaxOpen = int(getEnvDefaultInt64("DATABASE_MAX_OPEN", int64(cfg.Database.MaxOpen), 10))
	cfg.Database.MaxIdle = int(getEnvDefaultInt64("DATABASE_MAX_IDLE", int64(cfg.Database.MaxIdle), 10))

	// 日志配置
	cfg.Logger.Filename = getEnvDefaultString("LOGGER_FILENAME", cfg.Logger.Filename)
	cfg.Logger.Compress = getEnvDefaultBool("LOGGER_COMPRESS", cfg.Logger.Compress)
	cfg.Logger.MaxSizeMB = int(getEnvDefaultInt64("LOGGER_MAX_SIZE_MB", int64(cfg.Logger.MaxSizeMB), 10))
	cfg.Logger.MaxBackups = int(getEnvDefaultInt64("LOGGER_MAX_BACKUPS", int64(cfg.Logger.MaxBackups), 10))
	cfg.Logger.MaxDays = int(getEnvDefaultInt64("LOGGER_MAX_DAYS", int64(cfg.Logger.MaxDays), 30))
}

func initLogger(level, LogEncoding string, logCfg ...conf.LoggerConfig) {
	var err error

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	lv := zap.InfoLevel
	switch strings.ToLower(level) {
	case "debug":
		lv = zap.DebugLevel
	case "info":
		lv = zap.InfoLevel
	case "warn", "warning":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	default:
		lv = zap.InfoLevel
	}

	if len(logCfg) == 0 || logCfg[0].Filename == "" {
		cfg.Encoding = "console"
		if LogEncoding != "console" {
			cfg.Encoding = "json"
		}
		cfg.Level.SetLevel(lv)

		paths := []string{"stdout"}
		cfg.ErrorOutputPaths = paths
		cfg.OutputPaths = paths
		logger, err = cfg.Build()
		if err != nil {
			logger.Fatal("zap build", zap.Error(err))
		}
		return
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logCfg[0].Filename,
		MaxSize:    logCfg[0].MaxSizeMB, // megabytes
		MaxBackups: logCfg[0].MaxBackups,
		MaxAge:     logCfg[0].MaxDays, // days
		Compress:   logCfg[0].Compress,
	})

	enc := zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	if LogEncoding != "console" {
		enc = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	}
	core := zapcore.NewCore(
		enc,
		w,
		lv,
	)

	logger = zap.New(
		core,
		zap.AddCaller(),
		// zap.AddCallerSkip(1),
	)
}

func runServer() {
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

	if cfg.JWT.Secret == "" || cfg.JWT.Secret == "moredoc" {
		logger.Fatal("JWT.Secret", zap.String("安全风险提示", "JWT.Secret不能为空也不能为moredoc，请修改以保证安全性！！！"))
	}
	service.Run(cfg, logger)
}
