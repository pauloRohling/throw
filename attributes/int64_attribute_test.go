package attributes

import (
	"math"
	"testing"
)

func TestInt64Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int64
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         int64(20),
			expectedValue: "20",
		},
		"should return max int64 value": {
			key:           "key",
			value:         math.MaxInt64,
			expectedValue: "9223372036854775807",
		},
		"should return min int64 value": {
			key:           "key",
			value:         math.MinInt64,
			expectedValue: "-9223372036854775808",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt64(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
