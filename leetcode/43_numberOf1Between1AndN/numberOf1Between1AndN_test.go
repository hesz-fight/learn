package numberof1between1andn

import (
	"testing"
)

func Test_numberOf1Between1AndN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				n: 13,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOf1Between1AndN(tt.args.n); got != tt.want {
				t.Errorf("numberOf1Between1AndN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOf1Between1AndN2(t *testing.T) {
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
				n: 534,
			},
			want: 214,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOf1Between1AndN2(tt.args.n); got != tt.want {
				t.Errorf("numberOf1Between1AndN2() = %v, want %v", got, tt.want)
			}
		})
	}
}
