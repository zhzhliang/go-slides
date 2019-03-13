// +build OMIT
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error

	Value(key interface{}) interface{}
}