package livedata

type LiveDataInt struct {
	liveData *LiveData
}

func NewLiveDataInt(i int) *LiveDataInt {
	return &LiveDataInt{
		liveData: NewLiveData(i),
	}
}

func (l *LiveDataInt) ObserveForever(onChange func(int)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(int))
	})
}

func (l *LiveDataInt) Post(i int) {
	l.liveData.Post(i)
}

func (l *LiveDataInt) Get() int {
	return l.liveData.Get().(int)
}
