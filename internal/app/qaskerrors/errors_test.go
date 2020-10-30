package qaskerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrno_Error(t *testing.T) {
	tests := []struct {
		name string
		e    Errno
		want string
	}{
		{
			name: "Out of range",
			e:    1000,
			want: "errno: 1000",
		},
		{
			name: "Valid index",
			e:    1,
			want: "SQL duplicate entry",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("Errno.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrno_Is(t *testing.T) {
	type args struct {
		target error
	}
	tests := []struct {
		name string
		e    Errno
		args args
		want bool
	}{
		{
			name: "Is true #1",
			e:    EDUPENTRY,
			args: args{
				target: ErrUserExists,
			},
			want: true,
		},
		{
			name: "Is true #2",
			e:    EINVALIDMYSQLSYNTAX,
			args: args{
				target: ErrInvalidSQLSyntax,
			},
			want: true,
		},
		{
			name: "Is false #1",
			e:    EDUPENTRY,
			args: args{
				target: ErrUnknown,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Is(tt.args.target); got != tt.want {
				t.Errorf("Errno.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQaskErr_Error(t *testing.T) {
	type fields struct {
		Message string
		Code    uint16
		Err     error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Valid",
			fields: fields{
				Message: "Valid",
				Code:    100,
				Err:     nil,
			},
			want: "Valid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := QaskErr{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
				Err:     tt.fields.Err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("QaskErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQaskErr_Unwrap(t *testing.T) {
	type fields struct {
		Message string
		Code    uint16
		Err     error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid error",
			fields: fields{
				Message: "Valid",
				Code:    100,
				Err:     errors.New("Test error"),
			},
			wantErr: true,
		},
		{
			name: "Valid nil error",
			fields: fields{
				Message: "Valid",
				Code:    100,
				Err:     nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := QaskErr{
				Message: tt.fields.Message,
				Code:    tt.fields.Code,
				Err:     tt.fields.Err,
			}
			if err := e.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("QaskErr.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		s string
		c uint16
	}
	tests := []struct {
		name string
		args args
		want *QaskErr
	}{
		{
			name: "Valid",
			args: args{
				s: "Test error string",
				c: 100,
			},
			want: &QaskErr{
				Message: "Test error string",
				Code:    100,
				Err:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.s, tt.args.c)
			assert.NotNil(t, got)
			if got.Message != tt.want.Message ||
				got.Code != tt.want.Code ||
				got.Err != tt.want.Err {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
