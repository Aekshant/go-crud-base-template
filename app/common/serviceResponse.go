package common

type ServiceRepo[T any] struct {
	Data    T
	Err     error
	Success bool
}
