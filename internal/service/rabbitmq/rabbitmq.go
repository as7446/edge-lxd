package rabbitmq

import (
	"fmt"
	"github.com/as7446/lxd/internal/config"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	cfg     config.Config
}

func (r *RabbitMQ) NewRabbitMQ() (*RabbitMQ, error) {
	c := config.C()
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.RabbitMQ.Username,
		c.RabbitMQ.Password, c.RabbitMQ.IP, c.RabbitMQ.Port, c.RabbitMQ.VHost)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("连接rabbitmq失败: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("创建channel失败: %v", err)
	}
	rabbit := &RabbitMQ{conn: conn, channel: channel, cfg: c}
	if err := rabbit.declareExchange(); err != nil {
		rabbit.Close()
		return nil, err
	}
	return rabbit, nil
}

func (r *RabbitMQ) declareExchange() error {
	return r.channel.ExchangeDeclare(
		r.cfg.MetaData.Name, //本地hostname
		"topic",             //类型
		true,                // 持久化
		false,               // 非自动删除
		false,               // 非内部 Exchange
		false,               // 非阻塞
		nil,                 // 额外参数
	)
}

func (r *RabbitMQ) declareQueueAndBind(queueName, routingKey string) error {
	_, err := r.channel.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化
		false,     // 非独占
		false,     // 不自动删除
		false,     // 不等待服务器确认
		nil,       // 额外参数
	)
	if err != nil {
		return fmt.Errorf("声明队列失败: %v", err)
	}
	err = r.channel.QueueBind(
		queueName,
		routingKey,
		r.cfg.MetaData.Name,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("绑定队列失败: %v", err)
	}
	log.Printf("队列: %s 绑定到Exchanges %s, Routing key: %s", queueName, r.cfg.MetaData.Name, routingKey)
	return nil
}

type MessageHandler func(msg amqp.Delivery)

func (r *RabbitMQ) RegisterConsumer(queueName string, handler MessageHandler) error {
	err := r.declareQueueAndBind(queueName, queueName)
	if err != nil {
		log.Fatalf("队列注册失败: %v", err)
	}
	go func() {
		msgs, err := r.channel.Consume(
			queueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("队列消费失败: %v", err)
		}
		for msg := range msgs {
			handler(msg)
		}
	}()
	log.Printf("消费者已启动，监听队列: %v", queueName)
	return nil
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		_ = r.channel.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}
