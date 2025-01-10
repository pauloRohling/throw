package attributes

import (
	"math"
	"testing"
)

func TestInt16Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int16
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         int16(20),
			expectedValue: "20",
		},
		"should return max int16 value": {
			key:           "key",
			value:         math.MaxInt16,
			expectedValue: "32767",
		},
		"should return min int16 value": {
			key:           "key",
			value:         math.MinInt16,
			expectedValue: "-32768",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt16(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
