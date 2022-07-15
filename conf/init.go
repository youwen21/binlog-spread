package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type OptsOfSMTP struct {
	SmtpHost string `toml:"smtp_host" mapstructure:"smtp_host"`
	Port     int    `toml:"smtp_port" mapstructure:"smtp_port"`
	User     string `toml:"smtp_user" mapstructure:"smtp_user"`
	Password string `toml:"smtp_password" mapstructure:"smtp_password"`
}

type Mysql struct {
	Host     string `mapstructure:"db_host"`
	Username string `mapstructure:"db_username"`
	Password string `mapstructure:"db_password"`
	Port     int    `mapstructure:"db_port"`
	Database string `mapstructure:"db_database"`
	Charset  string `mapstructure:"db_charset"`
}

type BinlogMysql struct {
	ServerId int `mapstructure:"binlog_server_id"`

	Host     string `mapstructure:"binlog_db_host"`
	Username string `mapstructure:"binlog_db_username"`
	Password string `mapstructure:"binlog_db_password"`
	Port     int    `mapstructure:"binlog_db_port"`
	Database string `mapstructure:"binlog_db_database"`
	Charset  string `mapstructure:"binlog_db_charset"`

	Filter string `mapstructure:"binlog_filter"`
}

type CheckState struct {
	EnableCheckState string `mapstructure:"enable_check_state"`
}

type ModelStream struct {
	EnableModelStream        string `mapstructure:"enable_model_stream"`
	ModelStreamFlushRows     int    `mapstructure:"model_stream_flush_rows"`
	ModelStreamFlushDuration int64    `mapstructure:"model_stream_flush_duration"`
}

type AppConfig struct {
	AppEnv string `mapstructure:"app_env"`

	WebListen string `mapstructure:"web_listen"`

	DbType       string `mapstructure:"db_type"`
	SqliteFile   string `mapstructure:"sqlite_file"`
	DefaultMysql Mysql  `mapstructure:",squash"`

	BinlogMysql BinlogMysql `mapstructure:",squash"`

	Smtp OptsOfSMTP `mapstructure:",squash"`

	CheckState  `mapstructure:",squash"`
	ModelStream `mapstructure:",squash"`
}

var Config *AppConfig

func init() {
	flags := flag.NewFlagSet("anyName", flag.ExitOnError)
	cFile := flags.String("config", "config", "")
	// urfave/cli 程序 所以忽略第1，第2个参数
	_ = flags.Parse(os.Args[2:])

	configFile := strings.Replace(*cFile, filepath.Ext(*cFile), "", 1)
	//viper.AutomaticEnv()
	viper.SetConfigName(configFile) // name of config file (without extension)  config.type{env, yaml, toml , etc...}
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	//settings := viper.AllSettings()
	//fmt.Println(settings)

	err = viper.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("viper Unmarshal error: %w \n", err))
	}
}
