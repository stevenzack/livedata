package livedata

import "fmt"

const (
	chanName = "livedata/"
)

type Data struct {
	ID   string
	Data interface{}
	Slot func(interface{})
}

var (
	eventMap = make(map[string][]func(interface{}))
	dataMap  = make(map[string]interface{})
	add      = make(chan Data, 1)
	remove   = make(chan Data, 1)
	call     = make(chan Data, 1)
)

func init() {
	go func() {
		for {
			select {
			case data := <-add:
				fns := eventMap[data.ID]
				eventMap[data.ID] = append(fns, data.Slot)
			case data := <-remove:
				fns := eventMap[data.ID]
				for index, fn := range fns {
					if fmt.Sprint(fn) == fmt.Sprint(data.Slot) {
						eventMap[data.ID] = append(fns[:index], fns[index+1:]...)
						break
					}
				}
			case data := <-call:
				dataMap[data.ID] = data.Data
				fns := eventMap[data.ID]
				for _, fn := range fns {
					fn(data.Data)
				}
			}
		}
	}()
}

func Observe(id string, fn func(interface{})) {
	add <- Data{ID: id, Slot: fn}
}

func Dismiss(id string, fn func(interface{})) {
	remove <- Data{ID: id, Slot: fn}
}

func Set(id string, i interface{}) {
	call <- Data{ID: id, Data: i}
}

func Broadcast(id string) {
	call <- Data{ID: id, Data: dataMap[id]}
}

func Get(id string) interface{} {
	return dataMap[id]
}
