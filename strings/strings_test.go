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
		{"auto col narrow", "1 2 3", " ", -1, []string{"1", "2", "3"}},
		{"auto col wide", "1   2   3", " ", -1, []string{"1", "2", "3"}},
		{"one col narrow", "1 2 3", " ", 1, []string{"1 2 3"}},
		{"two col narrow", "1 2 3", " ", 2, []string{"1", "2 3"}},
		{"three col narrow", "1 2 3", " ", 3, []string{"1", "2", "3"}},
		{"empty", "", " ", -1, []string{}},
	} {
		t.Run(fmt.Sprintf("SplitNMerged() %s", tc.desc), func(t *testing.T) {
			got, want := SplitNMerged(tc.str, tc.sep, tc.cols), tc.want
			if !operators.EqualSlicesOfString(want, got) {
				t.Errorf("SplitNMerged(): want %v, got %v", want, got)
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
