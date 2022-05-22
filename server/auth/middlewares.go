package auth

func CheckAuth(token string) bool {
	if token == "" {
		return false
	}
	return true
}
