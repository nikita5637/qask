package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TestDB(t *testing.T) {
	type args struct {
		t           *testing.T
		databaseURL string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Invalid databaseURL",
			args: args{
				t:           t,
				databaseURL: "invalid",
			},
		},
		{
			name: "Invalid ping",
			args: args{
				t:           t,
				databaseURL: "root:12345678@tcp(256.255.0.0)/qask_test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { TestDB(tt.args.t, tt.args.databaseURL) })
		})
	}
}
