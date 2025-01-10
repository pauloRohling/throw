package attributes

import (
	"testing"
)

func TestStringAttribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         string
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         "value",
			expectedValue: "value",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewString(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
