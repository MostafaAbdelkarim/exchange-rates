package config

import "github.com/gofiber/fiber/v2/middleware/logger"

func GetLoggerConfig() logger.Config {

	var LoggerConfig = logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${latency} ${method} ${path} ${body}\n",
		TimeFormat: "yyyy-MM-dd",
		TimeZone:   "Africa/Cairo",
	}
	return LoggerConfig
}
