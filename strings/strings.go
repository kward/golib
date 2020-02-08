package strings

import (
	"strings"
	"unicode/utf8"
)

// SplitNMerged slices s into substrings separated by sep and returns a slice
// of the substrings between those separators. If sep is empty, SplitN splits
// after each UTF-8 sequence. Repeated sep will be merged. The count determines
// the number of substrings to return:
//
//  n > 0: at most n substrings; the last substring will be the unsplit remainder.
//  n == 0: the result is nil (zero substrings)
//  n < 0: all substrings
func SplitNMerged(s string, sep string, n int) []string {
	return genSplitMerged(s, sep, n)

	// split := strings.SplitN(s, sep, n)
	// merged := make([]string, 0, len(split))
	// for _, v := range split {
	// 	if v != "" {
	// 		merged = append(merged, v)
	// 	}
	// }
	// fmt.Printf("SplitNMerged() s: %q sep: %q n: %d merged: %q\n", s, sep, n, merged)
	// return merged
}

// explode splits s into a slice of UTF-8 strings,
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 sequences become correct encodings of U+FFFD.
//
// The following code was copied from go1.13.6.
func explode(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := 0; i < n-1; i++ {
		ch, size := utf8.DecodeRuneInString(s)
		a[i] = s[:size]
		s = s[size:]
		if ch == utf8.RuneError {
			a[i] = string(utf8.RuneError)
		}
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

// Generic split merged: splits after each instance of sep.
//
// The following code originated from go1.13.6, but includes modifications to
// merge separators.
func genSplitMerged(s, sep string, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explode(s, n)
	}
	if n < 0 {
		n = strings.Count(s, sep) + 1
	}

	a := make([]string, n)
	i := 0
	j := 0
	n--
	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		if m > 0 {
			a[j] = s[:m]
			j++
		}
		s = s[m+len(sep):]
		i++

		for strings.Index(s, sep) == 0 {
			s = s[len(sep):]
		}
	}
	a[j] = s
	return a[:j+1]
}

// Stretch a string to a given length by appending a character to the end.
func Stretch(s string, r rune, l int) string {
	// Special cases.
	if len(s) >= l {
		return s
	}
	return s + strings.Repeat(string(r), l-len(s))
}
