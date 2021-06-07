package conf

import (
	"flag"
	"time"

	"github.com/zxysilent/logs"

	"github.com/spf13/viper"
)

var (
	App       *Config           //运行配置实体
	defConfig = "./config.toml" //配置文件路径，方便测试
)

type Config struct {
	Mode string    `mapstructure:"mode"`
	Jwt  *jwtConf  `mapstructure:"jwt"`
	Http *httpConf `mapstructure:"http"`
	Orm  *ormConf  `mapstructure:"orm"`
	Db   *dbConf   `mapstructure:"db"`
}

// jwt config
type jwtConf struct {
	LoginKey     string `mapstructure:"login_key"`
	LoginPath    string `mapstructure:"login_path"`
	AuthKey      string `mapstructure:"auth_key"`
	AuthLifetime int    `mapstructure:"auth_lifetime"`
}

// http config
type httpConf struct {
	Address           string        `mapstructure:"address"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
}

// orm config
type ormConf struct {
	OrmIdle      int  `mapstructure:"orm_idle"`       //
	OrmOpen      int  `mapstructure:"orm_open"`       //
	OrmShow      bool `mapstructure:"orm_show"`       //显示sql
	OrmSync      bool `mapstructure:"orm_sync"`       //同步表结构
	OrmCacheUse  bool `mapstructure:"orm_cache_use"`  //是否使用缓存
	OrmCacheSize int  `mapstructure:"orm_cache_size"` //缓存数量
	OrmHijackLog bool `mapstructure:"orm_hijack_log"` //劫持日志
}

// db config
type dbConf struct {
	DbHost   string `mapstructure:"db_host"`   //数据库地址
	DbPort   int    `mapstructure:"db_port"`   //数据库端口
	DbUser   string `mapstructure:"db_user"`   //数据库账号
	DbPasswd string `mapstructure:"db_passwd"` //数据库密码
	DbName   string `mapstructure:"db_name"`   //数据库名称
	DbParams string `mapstructure:"db_params"` //数据库参数
}

func NewConfig() *Config {
	return &Config{
		Http: &httpConf{
			Address:           "",
			ReadTimeout:       20 * time.Second,
			WriteTimeout:      20 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			IdleTimeout:       10 * time.Second,
		},
	}
}

func Init() {
	var cfgFile string
	// 从启动命令中读取配置文件路径
	//go run main.go -configPath jjjkk(参数)
	flag.StringVar(&cfgFile, "configPath", defConfig, "path of mall config file.")
	flag.Parse()//解析命令行

	logs.Info("cfgFile=",cfgFile)
	//viper
	if cfgFile == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	} else {
		viper.SetConfigFile(cfgFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	cfg := NewConfig()//根据配置文件，配置实例化
	if err := viper.Unmarshal(cfg); err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	App = cfg
	//打印配置实例化，进行验证
	// logs.Info("*App=",*App)
	// logs.Info("App.Jwt=",App.Jwt)
}
func (app *Config) IsProd() bool {
	return app.Mode == "prod"
}
func (app *Config) IsDev() bool {
	return app.Mode == "dev"
}
