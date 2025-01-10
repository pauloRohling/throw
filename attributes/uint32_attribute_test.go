package attributes

import (
	"math"
	"testing"
)

func TestUint32Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         uint32
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         uint32(20),
			expectedValue: "20",
		},
		"should return max uint32 value": {
			key:           "key",
			value:         math.MaxUint32,
			expectedValue: "4294967295",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewUint32(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
