package urlJoin

import (
	"testing"
)

type TestCase struct {
	inputs         []string
	shouldError    bool
	expectedOutput string
	name           string
}

func TestJoin(t *testing.T) {
	testCases := []TestCase{
		{
			name:           "Test basic path joining",
			inputs:         []string{"www.google.com", "a", "b"},
			expectedOutput: "www.google.com/a/b",
		},
		{
			name:           "Test basic path joining with extra slashes",
			inputs:         []string{"www.google.com/", "/a/", "/b/"},
			expectedOutput: "www.google.com/a/b",
		},
	}

	for _, testCase := range testCases {
		result := Join(testCase.inputs...)

		if result != testCase.expectedOutput {
			t.Errorf("Join() with args: %v FAILED, expected '%v' but got value '%v'", testCase.inputs, testCase.expectedOutput, result)
		} else {
			t.Logf("Join() with args: %v PASSED, expected '%v' and got value '%v'", testCase.inputs, testCase.expectedOutput, result)
		}
	}
}
