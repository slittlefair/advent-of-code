package cipher

import "testing"

func TestCaesarCipher(t *testing.T) {
	type args struct {
		text     string
		shiftNum int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "applies Caesar Cipher to the given text shifted number of supplied times",
			args: args{
				text:     "qZmt-zixMtkozy-Ivhz-343",
				shiftNum: 343,
			},
			want: "vEry-encRypted-Name-343",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarCipher(tt.args.text, tt.args.shiftNum); got != tt.want {
				t.Errorf("CaesarCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
