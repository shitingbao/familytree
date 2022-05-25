package nsq

import (
	"encoding/json"
	"fmt"

	"github.com/nsqio/go-nsq"
)

const (
	tcpNsqDAddr = "127.0.0.1:4150"
)

type Config struct {
	Host string `toml:"host"`
}

type Nsqd struct {
	producer *nsq.Producer
}

func New(c *Config) *Nsqd {
	config := nsq.NewConfig()

	if c.Host == "" {
		c.Host = tcpNsqDAddr
	}

	p, err := nsq.NewProducer(c.Host, config)
	if err != nil {
		panic(err)
	}

	nsqd := &Nsqd{
		producer: p,
	}

	if err := nsqd.Ping(); err != nil {
		panic("[NSQ] Host:('" + c.Host + "') Ping failed...")
	}

	fmt.Printf("[NSQ] Host:('%s') started...", c.Host)

	return nsqd
}

func (n *Nsqd) Publish(topic string, arg interface{}) (err error) {
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}
	return n.producer.Publish(topic, b)
}

func (n *Nsqd) Ping() error {
	return n.producer.Ping()
}

func (n *Nsqd) Close() {
	n.producer.Stop()
}
