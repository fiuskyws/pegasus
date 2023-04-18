package helper

// ToPtr converts a given T to *T
func ToPtr[T any](v T) *T {
	return &v
}
