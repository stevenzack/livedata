package livedata

type LiveDataBool struct {
	liveData *LiveData
}

func NewLiveDataBool(b bool) *LiveDataBool {
	return &LiveDataBool{
		liveData: NewLiveData(b),
	}
}

func (l *LiveDataBool) ObserveForever(onChange func(bool)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(bool))
	})
}

func (l *LiveDataBool) Post(b bool) {
	l.liveData.Post(b)
}

func (l *LiveDataBool) Get() bool {
	return l.liveData.Get().(bool)
}
