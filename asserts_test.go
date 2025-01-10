package throw

import "reflect"

func assertEqual(t TestingT, expected, actual any) {
	if expected != actual {
		t.Errorf("expected: %v\n"+"got: %v", expected, actual)
	}
}

func assertIsType(t TestingT, expected, actual any) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("expected: %v\n"+"got: %v", expected, actual)
	}
}
