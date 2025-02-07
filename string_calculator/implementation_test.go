package string_calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want int64
	}{
		{
			"first test case",
			"",
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.expr)
			require.Equal(t, tt.want, result)
		})
	}
}
