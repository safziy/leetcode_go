package p350_p399

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				num: 14,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				num: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
