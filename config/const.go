package config

import "time"

const (
	JwtSignKey            = "jwt_secret"
	AccessSubject         = "AT"
	RefreshSubject        = "RT"
	AccessExpirationTime  = time.Hour * 24
	RefreshExpirationTime = time.Hour * 24 * 7
)
