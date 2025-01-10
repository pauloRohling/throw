package attributes

type TestingT interface {
	Errorf(format string, args ...interface{})
}

func assertEqual(t TestingT, expected, actual any) {
	if expected != actual {
		t.Errorf("expected: %v\n"+"got: %v", expected, actual)
	}
}
