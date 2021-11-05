package p1200_p1249

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case1",
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			wantAns: 4,
		},
		{
			name: "case2",
			args: args{
				arr:        []int{1, 3, 5, 7},
				difference: 1,
			},
			wantAns: 1,
		},
		{
			name: "case3",
			args: args{
				arr:        []int{1, 5, 7, 8, 5, 3, 4, 2, 1},
				difference: -2,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubsequence(tt.args.arr, tt.args.difference); gotAns != tt.wantAns {
				t.Errorf("longestSubsequence() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
