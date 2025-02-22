package maxgiftvalue

import "testing"

func Test_maxgiftvalue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxgiftvalue(tt.args.grid); got != tt.want {
				t.Errorf("maxgiftvalue() = %v, want %v", got, tt.want)
			}
		})
	}
}
