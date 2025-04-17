package minwindow

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "test",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "test",
			args: args{
				s: "aabca",
				t: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
