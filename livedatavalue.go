package livedata

func NewValue(v interface{}) *LiveData {
	return NewLiveData(v)
}
