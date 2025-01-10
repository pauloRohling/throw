package attributes

import (
	"testing"
)

func TestFloat64Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         float64
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         float64(20.32),
			expectedValue: "20.32",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewFloat64(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
