package middleware

func setInt(i int) *int {
	return &i
}

func setString(s string) *string {
	return &s
}

func setBool(t bool) *bool {
	return &t
}
