package validteStackSequences

import (
	"testing"
)

func Test_validteStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_1",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{2, 5, 4, 3, 1},
			},
			want: true,
		},
		{
			name: "test_2",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validteStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validteStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
