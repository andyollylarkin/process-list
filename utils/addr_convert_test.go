package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_coventHexIpToAddress(t *testing.T) {
	type args struct {
		hexAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid ip",
			args:    args{hexAddr: "0100007F"},
			want:    "127.0.0.1",
			wantErr: false,
		},
		{
			name:    "valid ip 2",
			args:    args{hexAddr: "0100A8C0"},
			want:    "192.168.0.1",
			wantErr: false,
		},
		{
			name:    "empty addr",
			args:    args{hexAddr: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CoventHexIpToAddress(tt.args.hexAddr)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_coventHexPortToAddress(t *testing.T) {
	type args struct {
		hexPort string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid port",
			args:    args{hexPort: "0016"},
			want:    "22",
			wantErr: false,
		},
		{
			name:    "valid port 2",
			args:    args{hexPort: "B806"},
			want:    "47110",
			wantErr: false,
		},
		{
			name:    "invalid port",
			args:    args{hexPort: "0"},
			want:    "0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CoventHexPortToAddress(tt.args.hexPort)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
