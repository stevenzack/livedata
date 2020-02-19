package livedata

import "github.com/StevenZack/pubsub"

type LiveData struct {
	value  interface{}
	server *pubsub.Server
	chanId string
}

func NewLiveData(v interface{}) *LiveData {
	server := pubsub.NewServer()
	server.Start()
	return &LiveData{
		value:  v,
		server: server,
		chanId: "livedata",
	}
}

func (l *LiveData) ObserveForever(onChange func(interface{})) {
	go func() {
		client := pubsub.NewClient(l.server)
		defer client.UnSub()
		client.Sub(l.chanId, onChange)
	}()
	onChange(l.value)
}

func (l *LiveData) Post(v interface{}) {
	l.value = v
	l.server.Pub(l.chanId, v)
}

func (l *LiveData) Get() interface{} {
	return l.value
}
