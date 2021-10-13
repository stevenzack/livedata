package livedata

import (
	"context"
	"sync/atomic"

	"github.com/StevenZack/pubsub"
)

type LiveData struct {
	value atomic.Value
	topic *pubsub.Topic
}

func NewLiveData(v interface{}) *LiveData {
	livedata := &LiveData{
		topic: pubsub.NewTopic(),
	}
	livedata.value.Store(v)
	return livedata
}

func (l *LiveData) Observe(lifecycleGoroutine func(), onChange func(v interface{})) {
	l.topic.Subscribe(lifecycleGoroutine, func() {
		onChange(l.value.Load())
	})
}

func (l *LiveData) ObserveCtx(ctx context.Context, onChange func(v interface{})) {
	l.topic.SubscribeCtx(ctx, func() {
		onChange(l.value.Load())
	})
}

func (l *LiveData) ObserveForever(onChange func(v interface{})) {
	l.ObserveCtx(context.Background(), onChange)
}

func (l *LiveData) Post(v interface{}) {
	l.value.Store(v)
	l.topic.Broadcast()
}

func (l *LiveData) Get() interface{} {
	return l.value.Load()
}
