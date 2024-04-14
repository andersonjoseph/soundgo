package model

import (
	"reflect"
	"testing"

	"github.com/andersonjoseph/soundgo/internal/shared"
)

type testHasher struct{}

func (h testHasher) Hash(p string) (string, error) {
	return p, nil
}

func (h testHasher) Compare(hp string, sp string) bool {
	return hp == sp
}

var hasher = testHasher{}

func TestPassword_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "password", fields: fields{value: "password"}, want: "password"},
		{name: "empty password", fields: fields{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Password{
				value: tt.fields.value,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Password.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPassword(t *testing.T) {
	type args struct {
		v      string
		hasher shared.SecretHasher
	}
	tests := []struct {
		name    string
		args    args
		want    Password
		wantErr bool
	}{
		{
			name:    "hash password",
			args:    args{v: "Str0ngP4s$w0Rd", hasher: hasher},
			want:    Password{value: "Str0ngP4s$w0Rd", hasher: hasher},
			wantErr: false,
		},
		{
			name:    "hash password",
			args:    args{v: "1234567890", hasher: hasher},
			want:    Password{value: "1234567890", hasher: hasher},
			wantErr: false,
		},
		{
			name:    "short password",
			args:    args{v: "shrt", hasher: hasher},
			want:    Password{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPassword(tt.args.v, tt.args.hasher)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
