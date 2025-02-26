package leftrotate

import "testing"

func Test_leftRotate(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				str: "abcdefg",
				n:   2,
			},
			want: "cdefgab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRotate(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("leftRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
