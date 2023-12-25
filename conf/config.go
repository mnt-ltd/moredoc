package conf

type LoggerConfig struct {
	Filename   string
	MaxSizeMB  int // 每份日志大小
	MaxBackups int // 日志保留多少份
	MaxDays    int // 保留多少天的日志
	Compress   bool
}

// Config app config
type Config struct {
	Level       string //
	LogEncoding string // log encoding, json or console
	Logger      LoggerConfig
	Port        int // listent port
	Database    Database
	JWT         JWT
}

type Database struct {
	DSN     string // data source name
	ShowSQL bool
	MaxIdle int
	MaxOpen int
	Prefix  string // table prefix, default is nd_
}

type JWT struct {
	Secret     string
	ExpireDays int64
}
