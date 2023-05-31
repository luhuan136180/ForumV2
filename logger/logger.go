package logger

import (
	"furumvv2/setting"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack" //用于切割
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var lg *zap.Logger

func Init(cfg *setting.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge) //日志输出位置
	encoder := getEncoder()                                                            //格式
	//自定义日志级别
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	//
	var core zapcore.Core
	//使用mode控制日志输出的位置
	if mode == "dev" {
		//进入开发模式，日志输出到终端
		consoleEnbcoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		core = zapcore.NewTee( //两个输出
			zapcore.NewCore(encoder, writeSyncer, l), //输入日志文件
			//输出日志到操作平台,把DebugLevel级别的日志输入中断
			zapcore.NewCore(consoleEnbcoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), //输出日志到操作平台
		)
	} else { //正常进入日志文件
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg = zap.New(core, zap.AddCaller())
	//通过zap.ReplaceGlobals(*)函数生成全局的logger
	//通过zap.L()拿去此logger***************************************
	//替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //设置时间的格式为人类可读
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //CapitalLevelEncoder将Level序列化为全大写字符串。
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//zap库中的原生模式
func GetLogWriter1() zapcore.WriteSyncer {
	//os包，创建文件test.log
	//file, _ := os.Create("leearn-Zap/test.log")

	//1可添加日志
	file, _ := os.OpenFile("leearn-Zap/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//指定日志传入的文件
	return zapcore.AddSync(file)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{ //返回一个接收切割的日志写入
		Filename:   filename,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		MaxSize:    maxSize,
	}
	return zapcore.AddSync(lumberJackLogger)
}

//Ginlogger接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
