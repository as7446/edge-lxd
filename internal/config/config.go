package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	BaseModel
	MetaData
	Server   ServerConfig   `mapstructure:"server"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
}
type BaseModel struct {
	ID       string    `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
type MetaData struct {
	Name        string            `json:"name"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type ServerConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Port     string `json:"port" mapstructure:"port"`
	Hostname string `json:"hostname" mapstructure:"hostname"`
}

type RabbitMQConfig struct {
	IP       string `json:"ip" mapstructure:"ip"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	VHost    string `json:"vhost" mapstructure:"vhost"`
	Queue    string `json:"queue" mapstructure:"queue"`
	Retry    int    `json:"retry" mapstructure:"retry"`
}

var cfg Config

func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	// 使用环境变量
	viper.SetEnvPrefix("LXD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("未找到配置文件，使用默认配置.")
		} else {
			log.Fatalf("读取配置文件错误: %v", err)
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("配置文件解析错误：%v", err)
	}
	cfg.BaseModel.CreateAt = time.Now()
	cfg.BaseModel.UpdateAt = time.Now()
	cfg.MetaData.Name = cfg.Server.Hostname
	if cfg.Server.Hostname == "" {
		name, err := os.Hostname()
		if err != nil {
			log.Fatalf("获取hostname失败: %v", err)
		}
		cfg.MetaData.Name = name
	}
	log.Println("配置文件加载完成.")
}
func C() Config {
	return cfg
}
