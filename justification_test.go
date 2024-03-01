package justification

import "testing"

func TestLeftJustifyLine(t *testing.T) {
	tests := []struct {
		words    []string
		width    int
		expected string
	}{
		{
			[]string{""},
			1,
			" ",
		},
		{
			[]string{"a"},
			1,
			"a",
		},
		{
			[]string{"a"},
			2,
			"a ",
		},
		{
			[]string{"a", "a"},
			3,
			"a a",
		},
		{
			[]string{"a", "a"},
			4,
			"a a ",
		},
		{
			[]string{"fabricator"},
			10,
			"fabricator",
		},
	}

	for tidx, test := range tests {
		result := leftJustifyLine(test.words, test.width)
		if len(result) != test.width {
			t.Errorf("Test[%d] width wrong: expected=%d got=%d",
				tidx, test.width, len(result))
		}

		if result != test.expected {
			t.Errorf("Test[%d] result wrong: expected=%q got=%q",
				tidx, test.expected, result)
		}
	}
}

func TestFullJustifyLine(t *testing.T) {
	tests := []struct {
		words    []string
		width    int
		expected string
	}{
		{
			[]string{""},
			1,
			" ",
		},
		{
			[]string{"a"},
			1,
			"a",
		},
		{
			[]string{"a"},
			2,
			"a ",
		},
		{
			[]string{"a", "a"},
			3,
			"a a",
		},
		{
			[]string{"a", "a"},
			4,
			"a  a",
		},
		{
			[]string{"fabricator"},
			10,
			"fabricator",
		},
	}

	for tidx, test := range tests {
		result := fullJustifyLine(test.words, test.width)
		if len(result) != test.width {
			t.Errorf("Test[%d] width wrong: expected=%d got=%d",
				tidx, test.width, len(result))
		}

		if result != test.expected {
			t.Errorf("Test[%d] result wrong: expected=%q got=%q",
				tidx, test.expected, result)
		}
	}
}

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		width    int
		expected []string
	}{
		{
			[]string{
				"This", "is", "an", "example", "of", "text",
				"justification.",
			},
			16,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			[]string{
				"What", "must", "be", "acknowledgment",
				"shall", "be",
			},
			16,
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			[]string{
				"Science", "is", "what", "we", "understand",
				"well", "enough", "to", "explain", "to", "a",
				"computer.", "Art", "is", "everything", "else",
				"we", "do",
			},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			[]string{
				"Science", "is", "what", "we", "understand",
				"well", "enough", "to", "explain", "to", "a",
				"computer.", "Art", "is", "everything", "else",
				"we", "do", "for", "human", "passion",
			},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do for human passion",
			},
		},
		{
			[]string{
				"Science", "is", "what", "we", "understand",
				"well", "enough", "to", "explain", "to", "a",
				"computer.", "Art", "is", "everything", "else",
				"we", "do", "for", "human", "joy",
			},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do for human joy    ",
			},
		},
	}

	for tidx, test := range tests {
		result := FullJustifiy(test.words, test.width)
		compareLines(t, tidx, test.expected, result)
	}
}

func compareLines(t *testing.T, testNo int, expected, result []string) {
	if len(result) != len(expected) {
		t.Errorf("Test[%d] length wrong: expected=%d got=%d",
			testNo, len(expected), len(result))
	}

	for eidx, line := range expected {
		if result[eidx] != line {
			t.Errorf(
				"Test[%d] line[%d] wrong: expected=%q got=%q",
				testNo, eidx, line, result[eidx])
		}
	}
}
