package model

import (
	"reflect"
	"testing"
)

func TestNewEmail(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Email
		wantErr bool
	}{
		{
			name:    "valid email",
			args:    args{v: "john.doe@example.com"},
			want:    Email{value: "john.doe@example.com"},
			wantErr: false,
		},
		{
			name:    "no domain",
			args:    args{v: "john.doe"},
			want:    Email{},
			wantErr: true,
		},
		{
			name:    "invalid characters",
			args:    args{v: "john.doe@ex!ample.com"},
			want:    Email{},
			wantErr: true,
		},
		{
			name:    "empty email",
			args:    args{v: ""},
			want:    Email{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "regular email",
			fields: fields{value: "hey@mail.com"},
			want:   "hey@mail.com",
		},
		{
			name:   "zero value",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Email{
				value: tt.fields.value,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Email.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
