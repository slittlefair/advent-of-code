package helpers

import "testing"

func TestIntSlicesAreEqual(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if the given decks aren't of equal length",
			args: args{
				slice1: []int{1, 2, 3, 4},
				slice2: []int{1, 2},
			},
			want: false,
		},
		{
			name: "returns false if the given decks aren't equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 6, 5, 7},
				slice2: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: false,
		},
		{
			name: "returns true if the given decks are equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				slice2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSlicesAreEqual(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
