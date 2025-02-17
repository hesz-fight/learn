package permutation

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				str: "abc",
			},
			want: []string{
				"abc", "acb", "bac", "bca", "cab", "cba",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
