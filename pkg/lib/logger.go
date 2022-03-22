package lib

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	*logrus.Logger
}

// NewLogger 构造日志服务
func NewLogger(logPath, module string) (*Logger, error) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	// 定义文件前缀和日志名称
	prefix := logPath + "/" + module
	latestLogFile := prefix + ".log"
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	//logger.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	//logger.SetFormatter(&logrus.JSONFormatter{})
	// If you wish to add the calling method as a field, instruct the logger via:
	//logger.SetReportCaller(true)

	logWriter, err := rotatelogs.New(
		prefix+"-%Y-%m-%d.log",                    // 生成实际文件名的模式
		rotatelogs.WithLinkName(latestLogFile),    // 生成日志软连接
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割周期
	)
	if err != nil {
		return nil, err
	}

	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: logWriter,
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.PanicLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
		},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))

	return &Logger{logger}, err
}
