package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "ADOBECODEBANC"
	t := "ABC"
	for idx := 0; idx < b.N; idx++ {
		minWindow(s, t)
	}
}
func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s = \"ADOBECODEBANC\", t = \"ABC\"",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
		{
			name: "s = \"a\", t = \"a\"",
			args: args{s: "a", t: "a"},
			want: "a",
		},
		{
			name: "s = \"a\", t = \"aa\"",
			args: args{s: "a", t: "aa"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
