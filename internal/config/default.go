package config

import "time"

var defaultConfig = map[string]any{
	"matching_repository.waiting_list_prefix": "waiting-list",
	"jwt.access_subject":                      "AT",
	"jwt.refresh_subject":                     "RT",
	"jwt.access_expiration_time":              "24h",
	"jwt.refresh_expiration_time":             "168h",
	"http_config.port":                        8080,
	"mysql.host":                              "localhost",
	"mysql.port":                              3306,
	"mysql.db_name":                           "gameapp_db",
	"redis.host":                              "localhost",
	"redis.port":                              6379,
	"redis.db":                                0,
	"matching_service.waiting_timeout":        "30s",
	"application.gracefull_shutdown_timeout":  time.Second * 5,
	"scheduler.match_players_interval":        time.Second * 3,
}
