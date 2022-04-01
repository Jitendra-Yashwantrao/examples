package reversestring

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test 1", args: args{s: ""}, want: ""},
		{name: "test 1", args: args{s: "hi"}, want: "hi"},
		{name: "test 2", args: args{s: "hi jitendra welcome to go"}, want: "go to welcome jitendra hi"},
		{name: "test 3", args: args{s: "the sky is blue"}, want: "blue is sky the"},
		{name: "test 2", args: args{s: "  hello world  "}, want: "world hello"},
		{name: "test 2", args: args{s: "a good   example"}, want: "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
