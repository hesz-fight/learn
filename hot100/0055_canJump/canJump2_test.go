package canjump

import "testing"

func Test_canJump2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump2(tt.args.nums); got != tt.want {
				t.Errorf("canJump2() = %v, want %v", got, tt.want)
			}
		})
	}
}
