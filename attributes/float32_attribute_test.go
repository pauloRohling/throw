package attributes

import (
	"testing"
)

func TestFloat32Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         float32
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         float32(20.32),
			expectedValue: "20.32",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewFloat32(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
