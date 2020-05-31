package apiserver

//Config ...
type Config struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
	DBIP       string `toml:"db_ip"`
	DBName     string `toml:"db_name"`
	DBDriver   string `toml:"db_driver"`
}

//NewConfig with default values
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
