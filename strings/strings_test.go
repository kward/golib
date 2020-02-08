package strings

import (
	"fmt"
	"testing"

	operators "github.com/kward/golib/operators"
)

func TestSplitNMerged(t *testing.T) {
	for _, tc := range []struct {
		desc string
		str  string
		sep  string
		cols int
		want []string
	}{
		{"auto col narrow", "1 22 333", " ", -1, []string{"1", "22", "333"}},
		{"auto col wide", "1   22   333", " ", -1, []string{"1", "22", "333"}},
		{"one col narrow", "1 22 333", " ", 1, []string{"1 22 333"}},
		{"one col wide", "1   22   333", " ", 1, []string{"1   22   333"}},
		{"two col narrow", "1 22 333", " ", 2, []string{"1", "22 333"}},
		{"two col wide", "1   22   333", " ", 2, []string{"1", "22   333"}},
		{"three col narrow", "1 22 333", " ", 3, []string{"1", "22", "333"}},
		{"three col wide", "1   22   333", " ", 3, []string{"1", "22", "333"}},
		{"empty", "", " ", -1, []string{""}},
	} {
		t.Run(fmt.Sprintf("SplitNMerged() %s", tc.desc), func(t *testing.T) {
			got, want := SplitNMerged(tc.str, tc.sep, tc.cols), tc.want
			if !operators.EqualSlicesOfString(want, got) {
				t.Errorf("SplitNMerged(): want %q, got %q", want, got)
			}
		})
	}
}

func TestStretch(t *testing.T) {
	for _, tc := range []struct {
		desc   string
		str    string
		sep    rune
		length int
		want   string
	}{
		{"long", "str1", ' ', 6, "str1  "},
		{"just right", "str2", ' ', 4, "str2"},
		{"short", "str3", ' ', 1, "str3"},
		{"empty", "", ' ', 5, "     "},
	} {
		t.Run(fmt.Sprintf("Stretch() %s", tc.desc), func(t *testing.T) {
			if got, want := Stretch(tc.str, tc.sep, tc.length), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
