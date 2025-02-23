package iicontinuesequence

import (
	"reflect"
	"testing"
)

func Test_continuesequence2(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{
				target: 9,
			},
			want: [][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := continuesequence2(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("continuesequence2() = %v, want %v", got, tt.want)
			}
		})
	}
}
