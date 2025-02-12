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
		{
			"two numbers",
			"1,2",
			3,
		},
		{
			"one number",
			"1",
			1,
		},
		{
			"3 numbers",
			"1,2,3",
			6,
		},
		{
			"4 numbers",
			"1,2,3,4",
			10,
		},
		{
			"strange scenario",
			"1\n2,3",
			6,
		},
		{
			"new delimiter",
			"//;\n1;2",
			3,
		},
		{
			"new delimiter with brackets",
			"//[***]\n1***2***3",
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.expr)
			require.Equal(t, tt.want, result)
		})
	}
}
