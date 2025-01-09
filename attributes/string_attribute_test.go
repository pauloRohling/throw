package attributes

import (
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, testCase.key, attribute.Key())
			assert.Equal(t, testCase.expectedValue, attribute.Value())
		})
	}
}
