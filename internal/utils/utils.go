package utils

func Map[T, U any](src []T, f func(T) U) []U {
	dest := make([]U, len(src))
	for i := range src {
		dest[i] = f(src[i])
	}
	return dest
}
