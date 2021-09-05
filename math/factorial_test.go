package factorial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorial(t *testing.T) {
	expected := int64(1)
	for i := int64(0); i <= int64(20); i++ {
		actual := Factorial(i).Int64()
		require.EqualValues(t, expected, actual)
		expected *= (i + 1)
	}
}
