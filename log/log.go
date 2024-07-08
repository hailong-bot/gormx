package log

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Magenta     = "\033[35m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
)

type LoggerWithLogrus struct {
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(config Config) logger.Interface {
	infoStr := "%s\n"
	warnStr := "%s\n"
	errStr := "%s\n"
	traceStr := "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr := "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr := "%s %s\n[%.3fms] [rows:%v] %s"

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + Reset
		errStr = Magenta + "%s\n" + Reset + Red + Reset
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}

	return &LoggerWithLogrus{
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

func Default() logger.Interface {
	return New(Config{
		SlowThreshold: 200 * time.Millisecond,
		Colorful:      true,
		LogLevel:      logger.Warn,
	})
}

func (l *LoggerWithLogrus) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *LoggerWithLogrus) Info(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Info {
		logrus.Infof(l.infoStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *LoggerWithLogrus) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Warn {
		logrus.Warnf(l.warnStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *LoggerWithLogrus) Error(ctx context.Context, s string, i ...interface{}) {
	if !strings.HasSuffix(s, "record not found") && l.LogLevel >= logger.Error {
		logrus.Errorf(l.errStr+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (l *LoggerWithLogrus) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				logrus.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logrus.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				}
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				logrus.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				logrus.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				logrus.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				logrus.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
