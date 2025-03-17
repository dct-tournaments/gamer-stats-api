package pointer

// ToString returns the value of the string pointer passed in or empty string if the pointer is nil.
func ToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// FromString returns a pointer to the passed string value, or nil, if passed value is a zero value.
func FromString(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}
