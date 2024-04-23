package skip_test

import (
	"testing"

	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-skip"
)

func TestAnyRunes(t *testing.T) {

	tests := []struct{
		CutSet []rune
		RuneScanner io.RuneScanner
		Expected string
	}{
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("Hello world! 🙂 wow.")),
			Expected:                                          "Hello world! 🙂 wow.",
		},

		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader(" Hello world! 🙂 wow.")),
			Expected:                         "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("  Hello world! 🙂 wow.")),
			Expected:                          "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("   Hello world! 🙂 wow.")),
			Expected:                           "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("    Hello world! 🙂 wow.")),
			Expected:                            "Hello world! 🙂 wow.",
		},

		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("\tHello world! 🙂 wow.")),
			Expected:                          "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("\t\tHello world! 🙂 wow.")),
			Expected:                            "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("\t\t\tHello world! 🙂 wow.")),
			Expected:                              "Hello world! 🙂 wow.",
		},
		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader("\t\t\t\tHello world! 🙂 wow.")),
			Expected:                                "Hello world! 🙂 wow.",
		},

		{
			CutSet: []rune{'\t', ' '},
			RuneScanner: utf8.NewRuneScanner(strings.NewReader(" \t\t   \t\t\t\t Hello world! 🙂 wow.")),
			Expected:                                         "Hello world! 🙂 wow.",
		},
	}

	testloop: for testNumber, test := range tests {

		var runescanner io.RuneScanner = test.RuneScanner

		err := skip.AnyRunes(runescanner, test.CutSet...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue testloop
		}

		{
			var buffer strings.Builder

			miniloop: for {
				r, size, err := runescanner.ReadRune()
				if 0 < size {
					buffer.WriteRune(r)
				}
				if io.EOF == err {
					break miniloop
				}
				if nil != err {
					t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
					t.Logf("ERROR: (%T) %s", err, err)
					continue testloop
				}
			}

			expected := test.Expected
			actual   := buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual (after skipping) 'content' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue testloop
			}
		}
	}
}
