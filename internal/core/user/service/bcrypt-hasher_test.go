package service_test

import (
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/service"
)

func hashPassword(t *testing.T, p string, hasher service.BcryptHasher) string {
	t.Helper()

	hp, err := hasher.Hash(p)

	if err != nil {
		t.Fatal(err)
	}

	return hp
}

func TestBcrypt_Compare(t *testing.T) {
	hasher := service.BcryptHasher{}

	type args struct {
		hp string
		sp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "correct password",
			args: args{hp: hashPassword(t, "1234567890", hasher), sp: "1234567890"},
			want: true,
		},
		{
			name: "wrong password",
			args: args{hp: hashPassword(t, "1234567890", hasher), sp: "123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasher.Compare(tt.args.hp, tt.args.sp); got != tt.want {
				t.Errorf("hasher.Compare(): %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
