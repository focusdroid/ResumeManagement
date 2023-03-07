package middleware

/**
 * @author: focusdroid
 * @description: 日志管理
 * @version: 1.0
 * @time：`2023-03-07 12:22:30`
**/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Loggoer() gin.HandlerFunc {
	logPath := "log/log"             // 日志存放目录
	linkName := "log/latest_log.log" // 软链接文件名
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src // 日志输出

	logger.SetLevel(logrus.DebugLevel)
	logWrite, _ := rotate.New( // 设置日志设置时间和软链接
		logPath+"%Y%m%d.log",
		rotate.WithMaxAge(7*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
		rotate.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{ // 使用lfshook设置日志级别写在logWrite这个函数中
		logrus.InfoLevel:  logWrite,
		logrus.FatalLevel: logWrite,
		logrus.DebugLevel: logWrite,
		logrus.ErrorLevel: logWrite,
		logrus.PanicLevel: logWrite,
		logrus.TraceLevel: logWrite,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{ // 格式化时间
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI
		referer := c.Request.Referer
		host := c.Request.Host

		entry := logger.WithFields(logrus.Fields{
			"HostTime":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Host":      host,
			"Method":    method,
			"Agent":     userAgent,
			"Path":      path,
			"Referer":   referer,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
