package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 全局变量，用来保存程序的所有配置信息
var Conf = &WebAppConfig{}

type WebAppConfig struct {
	AppConfig   AppConfig   `mapsturcture:"app"`
	LogConfig   LogConfig   `mapstructure:"log"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"app_port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type LogConfig struct {
	Level     string `mapstructure:"level"`
	Filename  string `mapstructure:"filename"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxBackup int    `mapstructure:"max_backup"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"mysql_host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"mysql_password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"mysql_port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"redis_host"`
	Password string `mapstructure:"redis_password"`
	Port     int    `mapstructure:"redis_port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	// viper.SetDefault("fileDir", "./")
	// viper.SetConfigFile("config.yaml") // 指定配置文件路径
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	// viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	// viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath("./")  // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		zap.L().Error("viper.ReadConfig() failed,err:%v\n", zap.Error(err))
		return
	}
	//把读取到的信息反序列化到conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal failed,err:%v\n", zap.Error(err))
	}
	//监听配置改变
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("配置文件修改了..")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper unmarshal failed,err:", zap.Error(err))
		}
	})
	return err
}
