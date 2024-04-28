package internal

import (
	"fmt"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

const (
	Console = "console"
	Zap     = "zap"
	GoZero  = "go-zero"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logType string
	switch kg.C.System.DbType {
	case "mysql":
		logType = kg.C.Mysql.LogType
	case "pgsql":
		logType = kg.C.Pgsql.LogType
	}
	if len(logType) == 0 || logType == Console {
		w.Writer.Printf(message, data...)
	} else if logType == Zap {
		kg.L.Info(fmt.Sprintf(message+"\n", data...))
	} else if logType == GoZero {
		logx.Debug(fmt.Sprintf(message+"\n", data...))
	}
}
