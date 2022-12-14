package conf

// Config app config
type Config struct {
	Level    string //
	Port     int    // listent port
	Database Database
}

type Database struct {
	DSN     string // data source name
	ShowSQL bool
	MaxIdle int
	MaxOpen int
	Prefix  string // table prefix, default is nd_
}
