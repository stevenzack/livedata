package livedata

type LiveData struct {
	value   interface{}
	channel chan interface{}
}

func NewLiveData(v interface{}) *LiveData {
	return &LiveData{
		value:   v,
		channel: make(chan interface{}),
	}
}

func (l *LiveData) ObserveForever(onChange func(interface{})) {
	go func() {
		for v := range l.channel {
			onChange(v)
		}
	}()
}

func (l *LiveData) Post(v interface{}) {
	l.value = v
	l.channel <- v
}

func (l *LiveData) Get() interface{} {
	return l.value
}
