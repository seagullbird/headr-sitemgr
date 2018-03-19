package config

var (
	// PORT is the serving port, this value is set during compile time
	PORT = "unset"
	// InitialTheme currently is set dead to gohugo-theme-ananke
	InitialTheme = "gohugo-theme-ananke"
	// Auth0Domain is the domain for my auth0 service
	Auth0Domain = "https://headr.auth0.com"
	// Auth0Audience is the identifier for Auth0 API 'headr-api'
	Auth0Audience = "https://api.headr.io"
)
