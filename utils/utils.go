package utils

func GetPointer[T any](obj T) *T {
	return &obj
}
