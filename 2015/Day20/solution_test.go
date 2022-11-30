package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deliverPresentsPart1(t *testing.T) {
	tests := []struct {
		name            string
		arg             int
		want            int
		errorAssertFunc assert.ErrorAssertionFunc
	}{
		{
			name:            "returns an error if no solution is found",
			arg:             0,
			want:            -1,
			errorAssertFunc: assert.Error,
		},
		{
			name:            "returns lowest house that reaches the target, advent of code example",
			arg:             150,
			want:            8,
			errorAssertFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deliverPresentsPart1(tt.arg)
			tt.errorAssertFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_deliverPresentsPart2(t *testing.T) {
	tests := []struct {
		name            string
		arg             int
		want            int
		errorAssertFunc assert.ErrorAssertionFunc
	}{
		{
			name:            "returns an error if no solution is found",
			arg:             0,
			want:            -1,
			errorAssertFunc: assert.Error,
		},
		{
			name:            "returns lowest house that reaches the target, advent of code example",
			arg:             150,
			want:            4,
			errorAssertFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deliverPresentsPart2(tt.arg)
			tt.errorAssertFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
