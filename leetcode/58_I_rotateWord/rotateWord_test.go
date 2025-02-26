package irotateword

import (
	"reflect"
	"testing"
)

func Test_rotateWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				words: []string{"11", "22"},
			},
			want: []string{"22", "11"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateWord(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
