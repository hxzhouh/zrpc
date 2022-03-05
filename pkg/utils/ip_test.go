package utils

import "testing"

func TestGetClientIp(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"test", "1223", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClientIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetClientIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
