package p150_p199

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				numerator:   1,
				denominator: 2,
			},
			want: "0.5",
		},
		{
			name: "case2",
			args: args{
				numerator:   2,
				denominator: 1,
			},
			want: "2",
		},
		{
			name: "case3",
			args: args{
				numerator:   2,
				denominator: 3,
			},
			want: "0.(6)",
		},
		{
			name: "case4",
			args: args{
				numerator:   4,
				denominator: 333,
			},
			want: "0.(012)",
		},
		{
			name: "case5",
			args: args{
				numerator:   1,
				denominator: 5,
			},
			want: "0.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
