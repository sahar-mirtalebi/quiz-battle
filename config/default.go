package config

var defaultConfig = map[string]any{
	"jwt.access_subject":          "AT",
	"jwt.refresh_subject":         "RT",
	"jwt.access_expiration_time":  "24h",
	"jwt.refresh_expiration_time": "168h",
	"http_config.port":            8080,
	"mysql.host":                  "localhost",
	"mysql.port":                  3306,
	"mysql.db_name":               "gameapp_db",
}
