package projectUtil

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logs = make(map[string]*logrus.Logger)

// GetFileLog 保存到指定日志文件
func GetFileLog(file ...string) *logrus.Logger {

	fileName := "default"
	if len(file) != 0 {
		fileName = file[0]
	}

	// 初始化日志
	if logs[fileName] == nil {
		var tempLog = logrus.New()
		tempLog.SetReportCaller(true)
		// 设置日志级别为xx以及以上
		tempLog.SetLevel(logrus.InfoLevel)
		logFile := path.Join("logs", fileName)
		// 本地
		//logFile := path.Join(".", "obData.log")
		writer, err := rotatelogs.New(
			logFile+".%Y%m%d%H%M",
			//rotatelogs.WithLinkName(logFile),                         // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(time.Duration(7*24)*time.Hour),     // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 日志切割时间间隔
		)
		if err != nil {
			tempLog.Errorf("config local file system logger error. %+v", err)
		}
		LogFileHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05", //时间格式化
		})
		tempLog.AddHook(LogFileHook)
		tempLog.SetOutput(os.Stdout)

		logs[fileName] = tempLog
	}

	return logs[fileName]
}
