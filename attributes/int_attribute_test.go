package attributes

import (
	"testing"
)

func TestIntAttribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         20,
			expectedValue: "20",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
