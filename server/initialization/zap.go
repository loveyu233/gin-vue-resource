package initialization

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// 1 定义一下logger使用的常量
var (
	maxSize    int
	maxAge     int
	maxBackups int
)

func InitLogger() {
	// 读取配置
	mode := viper.GetString("log.mode")          //开发模式
	maxSize = viper.GetInt("log.max-size")       //最大存储大小
	maxAge = viper.GetInt("log.max-age")         //最大存储时间
	maxBackups = viper.GetInt("log.max-backups") //#备份数量
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	now := time.Now().Format("2006-01-02")
	var cores [4]zapcore.Core
	if mode == "dev" {
		cores = [...]zapcore.Core{
			devModel(fmt.Sprintf("%s/%s/%s", "./log", now, "info.log"), infoPriority),
			devModel(fmt.Sprintf("%s/%s/%s", "./log", now, "debug.log"), debugPriority),
			devModel(fmt.Sprintf("%s/%s/%s", "./log", now, "warn.log"), warnPriority),
			devModel(fmt.Sprintf("%s/%s/%s", "./log", now, "errors.log"), errorPriority),
		}
	} else {
		cores = [...]zapcore.Core{
			zapLevelPrinf(fmt.Sprintf("%s/%s/%s", "./log", now, "info.log"), infoPriority),
			zapLevelPrinf(fmt.Sprintf("%s/%s/%s", "./log", now, "debug.log"), debugPriority),
			zapLevelPrinf(fmt.Sprintf("%s/%s/%s", "./log", now, "warn.log"), warnPriority),
			zapLevelPrinf(fmt.Sprintf("%s/%s/%s", "./log", now, "errors.log"), errorPriority),
		}
	}
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func zapLevelPrinf(filename string, leavel zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(zapEncode(), getLogWrite(filename), leavel)
}
func devModel(filename string, leavel zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(zapEncode(), zapcore.NewMultiWriteSyncer(getLogWrite(filename), zapcore.AddSync(os.Stdout)), leavel)
}

func zapEncode() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(ec)
}

func getLogWrite(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		// 日志文件的位置，也就是路径
		Filename: filename,
		// 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxSize: maxSize,
		// 保留旧文件的最大个数
		MaxBackups: maxBackups,
		// 保留旧文件的最大天数
		MaxAge: maxAge,
		// 是否压缩/归档旧文件
		Compress: true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
