package p600_p649

import "testing"

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				equation: "x=2",
			},
			want: "x=2",
		},
		{
			name: "case2",
			args: args{
				equation: "3x+16-2x=6x-2+x",
			},
			want: "x=3",
		},
		{
			name: "case3",
			args: args{
				equation: "-x+3x+18-2x=-6x-2+x",
			},
			want: "x=-4",
		},
		{
			name: "case4",
			args: args{
				equation: "x+5-3+x=6+x-2",
			},
			want: "x=2",
		},
		{
			name: "case5",
			args: args{
				equation: "x=x",
			},
			want: "Infinite solutions",
		},
		{
			name: "case6",
			args: args{
				equation: "2x=x",
			},
			want: "x=0",
		},
		{
			name: "case7",
			args: args{
				equation: "1+1=x",
			},
			want: "x=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveEquation(tt.args.equation); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
