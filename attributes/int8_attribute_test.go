package attributes

import (
	"math"
	"testing"
)

func TestInt8Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int8
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         int8(20),
			expectedValue: "20",
		},
		"should return max int8 value": {
			key:           "key",
			value:         math.MaxInt8,
			expectedValue: "127",
		},
		"should return min int8 value": {
			key:           "key",
			value:         math.MinInt8,
			expectedValue: "-128",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt8(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
