package iioccroncenum

import "testing"

func Test_occrOnceNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 5, 3, 3},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := occrOnceNum(tt.args.nums); got != tt.want {
				t.Errorf("occrOnceNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
