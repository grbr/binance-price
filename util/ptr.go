package util

func Ptr[T any](x T) *T {
	return &x
}
