package dbchgkinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_escape(t *testing.T) {
	testCases := []struct {
		name   string
		str    string
		wanted string
	}{
		{
			name:   `valid`,
			str:    "example quest",
			wanted: "example quest",
		},
		{
			name:   `invalid with \n`,
			str:    "example quest \n with \n",
			wanted: "example quest   with  ",
		},
		{
			name:   `invalid with slashes \`,
			str:    `example quest \" qwerty `,
			wanted: `example quest " qwerty `,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wanted, escape(tc.str))
		})
	}
}

func Test_parseAnswer(t *testing.T) {
	sr := New()
	tests := []struct {
		name         string
		answer       string
		wantedString string
		wantedError  bool
	}{
		{
			name:         "Valid",
			answer:       "\n    Ответ: Answer1",
			wantedString: "Answer1",
			wantedError:  false,
		},
		{
			name:         "Valid",
			answer:       "\n    Ответ: Answer1 SomethingText",
			wantedString: "Answer1 SomethingText",
			wantedError:  false,
		},
		{
			name:         "Not Valid",
			answer:       "\n  Ответ: Answer1",
			wantedString: "",
			wantedError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseAnswer(tt.answer, sr.validAnswerRegexp)
			if tt.wantedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			// assert.Equal(t, tt.wantedString, answer)
		})
	}
}
