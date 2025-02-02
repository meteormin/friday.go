package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadWithViper(t *testing.T) {
	type args struct {
		in      string
		appInfo App
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "test load with viper",
			args: args{
				in: "",
				appInfo: App{
					Name:    "Friday.go test",
					Version: "0.0.0",
					Mod:     "test",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadWithViper(tt.args.in, tt.args.appInfo)
			assert.NotNil(t, got)
			t.Log(*got)
		})
	}
}
