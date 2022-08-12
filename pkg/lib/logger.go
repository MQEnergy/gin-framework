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
func NewLogger(logPath, module string, debug bool) (*Logger, error) {
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
	if debug == true {
		logger.SetLevel(logrus.DebugLevel)
	}
	// 设置日志格式
	//logger.SetFormatter(&logrus.JSONFormatter{})
	// If you wish to add the calling method as a field, instruct the logger via:
	//logger.SetReportCaller(true)

	// 设置rotatelogs
	logWriter, err := rotatelogs.New(
		prefix+"-%Y-%m-%d.log",                    // 生成实际文件名的模式
		rotatelogs.WithLinkName(latestLogFile),    // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 设置最大保存时间(30天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)
	if err != nil {
		return nil, err
	}
	logger.AddHook(lfshook.NewHook(
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
