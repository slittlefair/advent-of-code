package main

import "testing"

func Test_runFights(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if there is an error parsing the Boss",
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		// don't have a good example, so just use the real advent of code questions since we know
		// we got the correct solution. It's too much to compare everything, so just make sure the
		// Boss is parsed and we will end up with the right Cost arrays
		{
			name: "runs fights against given boss",
			arg: []string{
				"Hit Points: 109",
				"Damage: 8",
				"Armor: 2",
			},
			want:    111,
			want1:   188,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := runFights(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("runFights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runFights() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("runFights() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
