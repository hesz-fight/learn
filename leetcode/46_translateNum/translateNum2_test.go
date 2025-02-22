package translatenum

import "testing"

func Test_translateNum2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "xxx",
			args: args{
				num: 12258,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum2(tt.args.num); got != tt.want {
				t.Errorf("translateNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
