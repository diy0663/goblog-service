package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	FileName   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxsize"`
	MaxBackups int    `yaml:"maxbackups"`
	MaxAges    int    `yaml:"maxages"`
	Compress   bool   `yaml:"compress"`
	Level      string `yaml:"-"`
}

var ZapLog *zap.Logger

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitZapLogger(logconfig *Log) (err error) {
	writeSyncer := getLogWriter(logconfig.FileName, logconfig.MaxSize, logconfig.MaxBackups, logconfig.MaxAges)
	encoder := getEncoder()
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(logconfig.Level))
	if err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// 给包的全局变量赋值
	ZapLog = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(ZapLog) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可 ??
	return nil

}
