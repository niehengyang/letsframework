package appconfig

type AppConfig struct {
	AppName   string
	RunMode   string
	PageSize  string
	JwtIssuer string

	Mysql  MysqlConfig
	Redis  RedisConfig
	Logger LoggerConfig
	Server ServerConfig
}

type MysqlConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP       string
	Port     string
	Password string
	DB       uint
}

type ServerConfig struct {
	Port int
}

type LoggerConfig struct {
	LogPath string
	LogFile string
}
