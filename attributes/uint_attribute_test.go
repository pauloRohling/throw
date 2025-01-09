package attributes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUintAttribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         uint
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
			attribute := NewUint(testCase.key, testCase.value)
			assert.Equal(t, testCase.key, attribute.Key())
			assert.Equal(t, testCase.expectedValue, attribute.Value())
		})
	}
}
