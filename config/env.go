package config

import (
	"os"
)

var (
	ServerPort = GetEnv("SERVER_PORT", "7005")
	// dev
	MongoUrl = GetEnv("MONGODB_URL", "mongodb://localhost:27017")
	// prod
	// MongoUrl        = GetEnv("MONGODB_URL", "mongodb://dexter:Cetnbcm88Cetnbcm88$@localhost:27017/?authSource=admin&readPreference=primary&authMechanism=SCRAM-SHA-256&appname=MongoDB%20Compass&ssl=false")
	MongoDatabase   = GetEnv("MONGODB_DATABASE", "BOTDB")
	JWTSecret       = GetEnv("JWT_SECRET", "R1BYcTVXVGNDU2JmWHVnZ1lnN0FKeGR3cU1RUSDFSJFHDKJH3K3HKHFDU45QXV4SDJONFZ3ckhwS1N0ZjNCYVkzZ0F4RVBSS1UzRENwRw==")
	JWTExpirationMs = GetEnv("JWT_EXPIRATION_MS", "86400000")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
