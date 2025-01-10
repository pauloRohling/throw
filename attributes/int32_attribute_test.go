package attributes

import (
	"math"
	"testing"
)

func TestInt32Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int32
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         int32(20),
			expectedValue: "20",
		},
		"should return max int32 value": {
			key:           "key",
			value:         math.MaxInt32,
			expectedValue: "2147483647",
		},
		"should return min int32 value": {
			key:           "key",
			value:         math.MinInt32,
			expectedValue: "-2147483648",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt32(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
