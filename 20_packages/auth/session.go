package auth

func extractSession() string{
	return "loggedin"
}

func GetSession() string {
	return extractSession()
}
