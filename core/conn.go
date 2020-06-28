package core

import (
	"context"
	"github.com/veypi/webds/message"
	"io"
	"time"
)

type (
	ErrorFunc      = message.FuncError
	DisconnectFunc = message.FuncBlank
	ConnectFunc    = message.FuncBlank
)

type ConnCfg interface {
	Validate()
	Ctx() context.Context
	Webds() Webds
	PingPeriod() time.Duration
	ReadTimeout() time.Duration
	ReadBufferSize() int64
}

type Connection interface {
	ID() string
	// 判断是个被动建立的连接还是主动建立的连接
	Passive() bool
	TargetID() string
	Level() int
	SetLevel(int)
	TargetUrl() string
	SetTargetID(string)
	SetTargetHost(string)
	SetTargetPort(uint)
	SetTargetPath(string)

	// 该方法仅主动方有效
	OnConnect(ConnectFunc) *message.Subscriber
	OnDisconnect(DisconnectFunc) *message.Subscriber
	OnError(ErrorFunc) *message.Subscriber
	FireOnError(err error)
	// 发送给对方节点广播消息
	Publisher(string) func(interface{})
	// 发送给对方节点非广播消息
	Echo(topic message.Topic, data interface{})
	// 通知对方自己订阅消息，并注册响应函数
	Subscribe(message.Topic, message.Func) *message.Subscriber
	Wait() error
	io.WriteCloser
	Type() int
	Alive() bool
}
