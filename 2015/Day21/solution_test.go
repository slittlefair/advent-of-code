package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runFights(t *testing.T) {
	tests := []struct {
		name               string
		arg                []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if there is an error parsing the Boss",
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
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
			want:               111,
			want1:              188,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := runFights(tt.arg)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}

}
