package attributes

import (
	"testing"
)

type AStructWithTags struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func TestJsonAttribute(t *testing.T) {
	testCases := map[string]struct {
		key           string
		value         any
		expectedValue string
	}{
		"should return a struct value": {
			key:           "key",
			value:         AStruct{Name: "value", Age: 10},
			expectedValue: "{\"Name\":\"value\",\"Age\":10}",
		},
		"should use tags to return a struct value": {
			key:           "key",
			value:         AStructWithTags{Name: "value", Age: 2},
			expectedValue: "{\"name\":\"value\",\"age\":2}",
		},
		"should omit empty fields": {
			key:           "key",
			value:         AStructWithTags{Name: "value", Age: 0},
			expectedValue: "{\"name\":\"value\"}",
		},
		"should return a struct value with a pointer": {
			key:           "key",
			value:         &AStruct{Name: "value", Age: 10},
			expectedValue: "{\"Name\":\"value\",\"Age\":10}",
		},
		"should return INVALID_JSON_VALUE when the value is not supported by json.Marshal": {
			key:           "key",
			value:         make(chan bool),
			expectedValue: "INVALID_JSON_VALUE",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			attribute := NewJson(testCase.key, testCase.value)
			assertEqual(t, testCase.key, attribute.Key())
			assertEqual(t, testCase.expectedValue, attribute.Value())
		})
	}
}
