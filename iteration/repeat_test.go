package iteration

import "testing"

func TestRepeat(t *testing.T) {
	count := 6
	repeated := Repeat("a", count)
	var expected string
	for i := 0; i < count; i++ {
		expected += "a"
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
