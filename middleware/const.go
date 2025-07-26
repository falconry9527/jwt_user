package middleware

var SkipAuthPaths = map[string]bool{
	"/api/users/getAll": true,
	"/api/users/login":  true,
}
