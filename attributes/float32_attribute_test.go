package attributes

import (
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, testCase.key, attribute.Key())
			assert.Equal(t, testCase.expectedValue, attribute.Value())
		})
	}
}
