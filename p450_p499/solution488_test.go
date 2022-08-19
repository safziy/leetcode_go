package p450_p499

import "testing"

func Test_findMinStep(t *testing.T) {
	type args struct {
		board string
		hand  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				board: "WRRBBW",
				hand:  "RB",
			},
			want: -1,
		},
		{
			name: "case2",
			args: args{
				board: "WWRRBBWW",
				hand:  "WRBRW",
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				board: "G",
				hand:  "GGGGG",
			},
			want: 2,
		},
		{
			name: "case4",
			args: args{
				board: "RBYYBBRRB",
				hand:  "YRBGB",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinStep(tt.args.board, tt.args.hand); got != tt.want {
				t.Errorf("findMinStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
