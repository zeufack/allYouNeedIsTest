package generic

import "testing"

func assertEqual[T comparable](t *testing.T, firstVal, secondVal T) {
	t.Helper()
	if firstVal != secondVal {
		t.Errorf("%+v and %+v should be equal", firstVal, secondVal)
	}
}

func assertNotEqual[T comparable](t *testing.T, firstVal, secondVal T) {
	t.Helper()
	if firstVal == secondVal {
		t.Errorf("didn't want %+v ", firstVal)
	}
}

func TestAssertFunction(t *testing.T) {
	t.Run("assert on integers", func(t *testing.T) {
		assertEqual(t, 1, 1)
		assertNotEqual(t, 1, 2)
	})

	t.Run("assert on string", func(t *testing.T) {
		assertEqual(t, "test", "nottest")
		assertNotEqual(t, "test", "notest")
	})

}
