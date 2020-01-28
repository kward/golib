// Basic max(x, y) routines that aren't provided by default.

package math

import "testing"

func TestMaxInt(t *testing.T) {
	for _, tt := range []struct {
		desc string
		a, b int
		want int
	}{
		{"equal", 2, 2, 2},
		{"max_is_a", 2, 1, 2},
		{"max_is_b", 1, 2, 2},
	} {
		a, b := tt.a, tt.b
		if got, want := MaxInt(tt.a, tt.b), tt.want; got != want {
			t.Errorf("%s: MaxInt(%d, %d) = %d, want %d", tt.desc, a, b, got, want)
		}
		// Ensure backward compatibility (until all references are removed).
		if got, want := Max(tt.a, tt.b), tt.want; got != want {
			t.Errorf("%s: Max(%d, %d) = %d, want %d", tt.desc, a, b, got, want)
		}
	}
}

func TestMinInt(t *testing.T) {
	for _, tt := range []struct {
		desc string
		a, b int
		want int
	}{
		{"equal", 1, 1, 1},
		{"min_is_a", 1, 2, 1},
		{"min_is_b", 2, 1, 1},
	} {
		a, b := tt.a, tt.b
		if got, want := MinInt(tt.a, tt.b), tt.want; got != want {
			t.Errorf("%s: MinInt(%d, %d) = %d, want %d", tt.desc, a, b, got, want)
		}
		// Ensure backward compatibility (until all references are removed).
		if got, want := Min(tt.a, tt.b), tt.want; got != want {
			t.Errorf("%s: Min(%d, %d) = %d, want %d", tt.desc, a, b, got, want)
		}
	}
}
