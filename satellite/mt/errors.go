package mt

import "testing"

func FailIfNotEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	err := Equals(expected, got)
	FailOnError(t, "values are not equal", err)
}

func ErrorIfNotEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	err := Equals(expected, got)
	ErrorOnError(t, "values are not equal", err)
}

func FailOnError(t *testing.T, errMsg string, err error) {
	t.Helper()
	if err == nil {
		return
	}
	if errMsg == "" {
		t.Fatalf("Unexpected error: %v", err)
	} else {
		t.Fatalf("%s: %v", errMsg, err)
	}
}

func ErrorOnError(t *testing.T, errMsg string, err error) {
	t.Helper()
	if err == nil {
		return
	}
	if errMsg == "" {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Errorf("%s: %v", errMsg, err)
	}
}

func FailWithoutError(t *testing.T, errMsg string, err error) {
	t.Helper()
	if err != nil {
		return
	}
	if errMsg != "" {
		t.Error(errMsg)
	} else {
		t.Error("expected error not returned")
	}
}
