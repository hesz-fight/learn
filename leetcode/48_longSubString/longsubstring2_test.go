package longsubstring

import "testing"

func Test_longsubstring2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				str: "abcabcbb",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longsubstring2(tt.args.str); got != tt.want {
				t.Errorf("longsubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}
