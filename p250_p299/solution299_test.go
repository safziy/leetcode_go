package main

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "case2",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
		{
			name: "case3",
			args: args{
				secret: "1",
				guess:  "0",
			},
			want: "0A0B",
		},
		{
			name: "case4",
			args: args{
				secret: "1",
				guess:  "1",
			},
			want: "1A0B",
		},
		{
			name: "case5",
			args: args{
				secret: "1234",
				guess:  "5678",
			},
			want: "0A0B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
