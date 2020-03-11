package livedata

type List struct {
	livedata *LiveData
}

func NewList(l []interface{}) *List {
	return &List{
		livedata: NewLiveData(l),
	}
}

func (l *List) ObserveForever(onChange func([]interface{})) {
	l.livedata.ObserveForever(func(v interface{}) {
		onChange(v.([]interface{}))
	})
}

func (l *List) Post(list []interface{}) {
	l.livedata.Post(list)
}

func (l *List) GetAll() []interface{} {
	return l.livedata.Get().([]interface{})
}

func (l *List) Add(v interface{}) {
	l.livedata.Post(append(l.GetAll(), v))
}

func (l *List) Remove(i int) {
	list := l.GetAll()
	if i >= len(list) {
		return
	}
	l.livedata.Post(append(list[:i], list[i+1:]...))
}

func (l *List) Get(i int) interface{} {
	list := l.GetAll()
	if i >= len(list) {
		return nil
	}
	return list[i]
}
