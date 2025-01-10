package attributes

import (
	"math"
	"testing"
)

func TestUint64Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         uint64
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         uint64(20),
			expectedValue: "20",
		},
		"should return max uint64 value": {
			key:           "key",
			value:         math.MaxUint64,
			expectedValue: "18446744073709551615",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewUint64(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
