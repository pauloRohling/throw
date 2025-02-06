package throw

import "reflect"

func assertEqual(t TestingT, expected, actual any) {
	if expected != actual {
		t.Errorf("expected: %v, "+"got: %v", expected, actual)
	}
}

func assertNotNil(t TestingT, actual any) {
	if actual == nil {
		t.Errorf("Expected value not to be nil")
	}
}

func assertIsType(t TestingT, expected, actual any) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("expected: %v, "+"got: %v", expected, actual)
	}
}
