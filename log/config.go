package log

import (
	"time"

	"gorm.io/gorm/logger"
)

type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}
