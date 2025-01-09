package attributes

import (
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, testCase.key, attribute.Key())
			assert.Equal(t, testCase.expectedValue, attribute.Value())
		})
	}
}
