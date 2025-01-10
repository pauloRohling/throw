package attributes

import (
	"testing"
)

type AStruct struct {
	Name string
	Age  int
}

func TestAnyAttribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         any
		expectedValue string
	}{
		"should return the key and value": {
			key:           "key",
			value:         "value",
			expectedValue: "value",
		},
		"should return an integer value": {
			key:           "key",
			value:         20,
			expectedValue: "20",
		},
		"should return a float value": {
			key:           "key",
			value:         20.5,
			expectedValue: "20.5",
		},
		"should return a boolean value": {
			key:           "key",
			value:         true,
			expectedValue: "true",
		},
		"should return a struct value": {
			key:           "key",
			value:         AStruct{Name: "Lorem Ipsum", Age: 20},
			expectedValue: "{Name:Lorem Ipsum Age:20}",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewAny(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
