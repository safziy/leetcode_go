package p450_p499

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "5F3Z-2e-9-w",
				k: 4,
			},
			want: "5F3Z-2E9W",
		},
		{
			name: "case2",
			args: args{
				s: "2-5g-3-J",
				k: 2,
			},
			want: "2-5G-3J",
		},
		{
			name: "case3",
			args: args{
				s: "2-4A0r7-4k",
				k: 3,
			},
			want: "24-A0R-74K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
