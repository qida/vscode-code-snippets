/*
 * @Author: qida
 * @LastEditors: qida
 */
package sms

import "testing"

func TestCheckRegexMobile(t *testing.T) {
	type args struct {
		mobile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "16777259176",
			args: args{
				mobile: "16777259176",
			},
			wantErr: false,
		},
		{
			name: "1677725917",
			args: args{
				mobile: "1677725917",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckRegexMobile(tt.args.mobile); (err != nil) != tt.wantErr {
				t.Errorf("CheckRegexMobile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
