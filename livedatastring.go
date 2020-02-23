package livedata

type String struct {
	liveData *LiveData
}

func NewString(s string) *String {
	return &String{
		liveData: NewLiveData(s),
	}
}

func (l *String) ObserveForever(onChange func(string)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(string))
	})
}

func (l *String) Post(s string) {
	l.liveData.Post(s)
}

func (l *String) Get() string {
	return l.liveData.Get().(string)
}
