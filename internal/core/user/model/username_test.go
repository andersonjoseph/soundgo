package model

import (
	"reflect"
	"testing"
)

func TestUsername_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "lowercase username",
			fields: fields{value: "andersonjoseph"},
			want:   "andersonjoseph",
		},
		{
			name:   "mixed case username",
			fields: fields{value: "andersonjoseph"},
			want:   "andersonjoseph",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Username{
				value: tt.fields.value,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("Username.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUsername(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Username
		wantErr bool
	}{
		{
			name:    "Valid username",
			args:    args{v: "andersonjoseph"},
			want:    Username{value: "andersonjoseph"},
			wantErr: false,
		},
		{
			name:    "Valid username with camelCase",
			args:    args{v: "andersonJoseph"},
			want:    Username{value: "andersonJoseph"},
			wantErr: false,
		},
		{
			name:    "Empty username",
			args:    args{v: ""},
			want:    Username{},
			wantErr: true,
		},
		{
			name:    "Valid username with numbers",
			args:    args{v: "anderson123"},
			want:    Username{value: "anderson123"},
			wantErr: false,
		},
		{
			name:    "Valid username with underscore",
			args:    args{v: "anderson_joseph"},
			want:    Username{value: "anderson_joseph"},
			wantErr: false,
		},
		{
			name:    "Username with spaces",
			args:    args{v: "anderson joseph"},
			want:    Username{},
			wantErr: true,
		},
		{
			name:    "Username with special characters",
			args:    args{v: "anderson*joseph"},
			want:    Username{},
			wantErr: true,
		},
		{
			name:    "Username with mixed case and numbers",
			args:    args{v: "AndersonJoseph123"},
			want:    Username{value: "AndersonJoseph123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUsername(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
