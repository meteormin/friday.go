package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	type args struct {
		config []ZapLoggerConfig
	}
	tests := []struct {
		name string
		args args
		want *zap.SugaredLogger
	}{
		{
			name: "test new zap logger",
			args: args{
				config: []ZapLoggerConfig{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewZapLogger(tt.args.config...)
			assert.NotNil(t, got)
			got.Info("test")
		})
	}
}
