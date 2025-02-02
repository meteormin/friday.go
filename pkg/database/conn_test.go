package database

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		{
			name: "test new gorm.DB",
			args: args{
				cfg: Config{
					Name:   "test",
					Debug:  true,
					Logger: LoggerConfig{},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cfg)
			assert.NotNil(t, got)
		})
	}
}
