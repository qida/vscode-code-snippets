package sms

import (
	"testing"
)

func TestInitUMP(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitUMP(); (err != nil) != tt.wantErr {
				t.Errorf("InitUMP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		mobile string
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, err := Send(tt.args.mobile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("Send() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
