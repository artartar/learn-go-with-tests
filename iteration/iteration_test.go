package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but %q", expected, repeated)
	}
}
# Just getting myself to write some kind of code, even the function below
func TestVarRepeat(t *testing.T, x int) {
	x = 10
	repeated := Repeat("a", x)
	expected := "a" * x

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
