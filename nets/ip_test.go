//from https://github.com/freshcn/qqwry

package nets

import "testing"

func TestQQwry_Find(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		q    *QQwry
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			q:    NewQQwry("qqwry.dat"),
			args: args{ip: "39.78.34.61"},
			want: "山东省临沂市联通",
		},
		{
			name: "test2",
			q:    NewQQwry("qqwry.dat"),
			args: args{ip: "39.78.34.62"},
			want: "山东省临沂市联通",
		},
		{
			name: "test3",
			q:    NewQQwry("qqwry.dat"),
			args: args{ip: "39.78.34.63"},
			want: "山东省临沂市联通",
		},
		{
			name: "test4",
			q:    NewQQwry("qqwry.dat"),
			args: args{ip: "39.78.34.64"},
			want: "山东省临沂市联通",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Find(tt.args.ip); got != tt.want {
				t.Errorf("QQwry.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
