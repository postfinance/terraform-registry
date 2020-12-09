// Package logger returns a ready-to-use logger
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a logger
func New(debug bool) (*zap.SugaredLogger, error) {
	atom := zap.NewAtomicLevel()
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = atom
	zapConfig.Encoding = "console"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if debug {
		atom.SetLevel(zap.DebugLevel)
	}

	l, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return l.Sugar(), nil
}
