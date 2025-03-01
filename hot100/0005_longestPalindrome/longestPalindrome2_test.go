package longestpalindrome

import "testing"

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "abbba",
			},
			want: "abbba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
