package logg

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type Log struct {
	LogLevel logger.LogLevel
}

// LogMode log mode
func (l *Log) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l Log) Info(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
}

func (l Log) Warn(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
}

func (l Log) Error(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
}

func (l Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
}
