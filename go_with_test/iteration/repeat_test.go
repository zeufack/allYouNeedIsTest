package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("repeate character 5 time", func(t *testing.T) {
		repeat := Repeat("a", 5)
		expect := "aaaaa"

		if repeat != expect {
			t.Errorf("expect %q, bu got %q", expect, repeat)
		}
	})

	t.Run("repeate character one time", func(t *testing.T) {
		repeat := Repeat("b", 1)
		expect := "b"

		if repeat != expect {
			t.Errorf("expect %q, bu got %q", expect, repeat)
		}
	})
}

func BenchmarkRepaeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
