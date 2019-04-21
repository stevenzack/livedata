package livedata

import "github.com/StevenZack/pubsub"

const (
	chanName = "livedata/"
)

var (
	dataMap     = make(map[string]interface{})
	dataMapChan = make(chan Data, 1)
	ps          = pubsub.NewPubSub()
)

type Data struct {
	ID   string
	Data interface{}
}

func init() {
	go func() {
		for i := range dataMapChan {
			dataMap[i.ID] = i.Data
		}
	}()
}
func Observe(id string, f func(interface{})) {
	go func() {
		ps.Sub(chanName+id, f)
	}()
}

func Set(id string, i interface{}) {
	dataMapChan <- Data{ID: id, Data: i}
	pubsub.Pub(chanName+id, i)
}

func Get(id string) interface{} {
	return dataMap[id]
}
