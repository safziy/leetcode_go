package p1400_p1449

import "testing"

func Test_reformat(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "case1",
			args: "a0b1c2",
			want: "a0b1c2",
		},
		{
			name: "case2",
			args: "leetcode",
			want: "",
		},
		{
			name: "case3",
			args: "1229857369",
			want: "",
		},
		{
			name: "case4",
			args: "covid2019",
			want: "c0v9d2o1i",
		},
		{
			name: "case5",
			args: "ab123",
			want: "2b1a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.args); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
