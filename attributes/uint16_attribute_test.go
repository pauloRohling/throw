package attributes

import (
	"math"
	"testing"
)

func TestUint16Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         uint16
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         uint16(20),
			expectedValue: "20",
		},
		"should return max uint16 value": {
			key:           "key",
			value:         math.MaxUint16,
			expectedValue: "65535",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewUint16(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
