package printminnumber

import "testing"

func TestPrintMinNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test2",
			args: args{
				numbers: []int{3, 30, 34, 5, 9},
			},
			want: "3033459",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintMinNumber(tt.args.numbers); got != tt.want {
				t.Errorf("PrintMinNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
