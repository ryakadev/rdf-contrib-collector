package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a new instance of *zap.Logger.
//
// It takes a level of zapcore.Level as a parameter.
// It returns a pointer to *zap.Logger.
func New(level zapcore.Level) *zap.Logger {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()
	atom.SetLevel(level)

	return logger

}
