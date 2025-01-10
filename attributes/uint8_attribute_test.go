package attributes

import (
	"math"
	"testing"
)

func TestUint8Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         uint8
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         uint8(20),
			expectedValue: "20",
		},
		"should return max uint8 value": {
			key:           "key",
			value:         math.MaxUint8,
			expectedValue: "255",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewUint8(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
