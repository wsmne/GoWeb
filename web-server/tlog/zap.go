package tlog

import (
	"log/slog"
	"os"
	"path"

	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZlogHandler(dir string) slog.Handler {
	logger := newZapLogger(dir)
	handler := slogzap.Option{Level: slog.LevelDebug, AddSource: true, Logger: logger}.NewZapHandler()
	return NewHandler(handler)
}

func newZapLogger(dir string) *zap.Logger {
	if dir == "" {
		dir = path.Join("logs", "kernel")
	}
	var coreArr []zapcore.Core

	format := "2006-01-02 15:04:05.000"
	encoderConfig := zap.NewProductionEncoderConfig()              // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(format) // 指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder        // 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder        // 显示完整文件路径
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 日志级别
	errLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.InfoLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})

	// info文件writeSyncer
	debugFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(dir, "debug.log"), // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    256,                         // 文件大小限制,单位MB
		MaxBackups: 2,                           // 最大保留日志文件数量
		MaxAge:     2,                           // 日志文件保留天数
		Compress:   false,                       // 是否压缩处理
	})
	debugFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(debugFileWriteSyncer, zapcore.AddSync(os.Stdout)), debugLevel) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(dir, "info.log"), // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    256,                        // 文件大小限制,单位MB
		MaxBackups: 5,                          // 最大保留日志文件数量
		MaxAge:     1,                          // 日志文件保留天数
		Compress:   false,                      // 是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), infoLevel) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(dir, "error.log"), // 日志文件存放目录
		MaxSize:    256,                         // 文件大小限制,单位MB
		MaxBackups: 5,                           // 最大保留日志文件数量
		MaxAge:     1,                           // 日志文件保留天数
		Compress:   false,                       // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), errLevel) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, debugFileCore)
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	return zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) // zap.AddCaller()为显示文件名和行号，可省略
}
