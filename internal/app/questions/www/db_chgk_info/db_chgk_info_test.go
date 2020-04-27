package db_chgk_info

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
