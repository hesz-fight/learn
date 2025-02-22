package firstuniqchar

import "testing"

func Test_uniqueChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test",
			args: args{
				str: "abaccdeff",
			},
			want: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueChar(tt.args.str); got != tt.want {
				t.Errorf("uniqueChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
