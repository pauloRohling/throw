package attributes

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestInt8Attribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         int8
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         int8(20),
			expectedValue: "20",
		},
		"should return max int8 value": {
			key:           "key",
			value:         math.MaxInt8,
			expectedValue: "127",
		},
		"should return min int8 value": {
			key:           "key",
			value:         math.MinInt8,
			expectedValue: "-128",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewInt8(testCase.key, testCase.value)
			assert.Equal(t, testCase.key, attribute.Key())
			assert.Equal(t, testCase.expectedValue, attribute.Value())
		})
	}
}
