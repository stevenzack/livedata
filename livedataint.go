package livedata

type Int struct {
	liveData *LiveData
}

func NewInt(i int) *Int {
	return &Int{
		liveData: NewLiveData(i),
	}
}

func (l *Int) ObserveForever(onChange func(int)) {
	l.liveData.ObserveForever(func(v interface{}) {
		onChange(v.(int))
	})
}

func (l *Int) Post(i int) {
	l.liveData.Post(i)
}

func (l *Int) Get() int {
	return l.liveData.Get().(int)
}
