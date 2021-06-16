/*
 * @Author: qida
 * @LastEditors: qida
 */
package stringx

import "testing"

func TestGetKeysString(t *testing.T) {
	type args struct {
		key_str string
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
		wantPy     string
		wantHan    string
	}{
		{name: "1", args: args{key_str: "string"}, wantNumber: 0, wantPy: "STRING", wantHan: ""},
		{name: "2", args: args{key_str: "123"}, wantNumber: 123, wantPy: "", wantHan: ""},
		{name: "3", args: args{key_str: "好人"}, wantNumber: 0, wantPy: "", wantHan: "好人"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumber, gotPy, gotHan := GetKeysString(tt.args.key_str)
			if gotNumber != tt.wantNumber {
				t.Errorf("GetKeysString() gotNumber = %v, want %v", gotNumber, tt.wantNumber)
			}
			if gotPy != tt.wantPy {
				t.Errorf("GetKeysString() gotPy = %v, want %v", gotPy, tt.wantPy)
			}
			if gotHan != tt.wantHan {
				t.Errorf("GetKeysString() gotHan = %v, want %v", gotHan, tt.wantHan)
			}
		})
	}
}
