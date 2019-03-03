package api

const (
	// Version of Rikka
	Version = "0.5.0"

	// FormKeyFile is file field name when upload image
	FormKeyFile = "uploadFile"
	// FormKeyPWD is password field name when upload image
	FormKeyPWD = "password"
	// FormKeyFrom is from field name when upload image
	FormKeyFrom = "from"

	// FromWebsite is a value of FromKeyFrom, means request comes from website
	FromWebsite = "website"
	// FromAPI is a value of FromKeyFrom, means request comes from REST API
	FromAPI = "api"
)
