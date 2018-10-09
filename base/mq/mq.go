package mq

import (
	"sync"

	. "github.com/rockin0098/flash/base/logger"
)

type MQ interface {
	CreateChannel(channel string)
	Put(channel string, data interface{})
	Get(channel string) interface{}
}

func NewMQ() MQ {

	// 单机版
	return &MQInternal{
		channelMap: &sync.Map{},
	}
}

type MQInternal struct {
	channelMap *sync.Map
}

func (m *MQInternal) CreateChannel(channel string) {

}

func (m *MQInternal) Put(channel string, data interface{}) {
	ch, ok := m.channelMap.Load(channel)
	if !ok {
		Log.Errorf("channel=[%v] not found", channel)
		return
	}

	ch.(chan interface{}) <- data
}

func (m *MQInternal) Get(channel string) interface{} {
	ch, ok := m.channelMap.Load(channel)
	if !ok {
		Log.Errorf("channel=[%v] not found", channel)
		return nil
	}

	data := <-ch.(chan interface{})

	return data
}
