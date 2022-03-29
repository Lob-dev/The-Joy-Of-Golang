package util

func IsNilThrow(condition bool, message string) {
	if condition {
		panic(message)
	}
}
