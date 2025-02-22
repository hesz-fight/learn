package uglynum

import "testing"

func Test_uglyNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 10,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uglyNum(tt.args.n); got != tt.want {
				t.Errorf("uglyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
