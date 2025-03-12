package main

import (
	"fmt"
	"github.com/as7446/lxd/internal/config"
	"github.com/as7446/lxd/internal/service/rabbitmq"
	"github.com/spf13/cobra"
	"time"
)

var configPath string
var rootCmd = &cobra.Command{
	Use:   "lxd-api",
	Short: "LXD API 服务",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig(configPath)
		rabbit := rabbitmq.RabbitMQ{}
		mq, err := rabbit.NewRabbitMQ()
		if err != nil {
			fmt.Printf("new mq: %v", err)
			return
		}
		rabbitmq.RegisterQueueHandlers(mq)
		fmt.Println("服务启动中...")
		for {
			time.Sleep(time.Second * 10)
		}
	},
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "../conf/app.yml", "指定配置文件路径")
	rootCmd.Execute()
}
