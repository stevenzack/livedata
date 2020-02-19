package livedata

type LiveDataString struct {
	liveData *LiveData
}

func NewLiveDataString(s string) *LiveDataString {
	return &LiveDataString{
		liveData: NewLiveData(s),
	}
}

func (l *LiveDataString) ObserveForever(onChange func(string)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(string))
	})
}

func (l *LiveDataString) Post(s string) {
	l.liveData.Post(s)
}

func (l *LiveDataString) Get() string {
	return l.liveData.Get().(string)
}
