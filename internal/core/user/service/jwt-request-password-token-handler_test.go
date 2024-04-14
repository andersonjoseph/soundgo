package service_test

import (
	"testing"

	"github.com/andersonjoseph/soundgo/internal/core/user/service"
)

func TestJWTPasswordTokenGenerator_Generate(t *testing.T) {
	generator := service.NewJWTPasswordTokenGenerator([]byte("secret"))

	type args struct {
		userID int
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "generate a token",
			args: args{
				userID: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generator.Generate(tt.args.userID)
			tt.want = got

			if (err != nil) != tt.wantErr {
				t.Errorf("generator.Generate(): %s = %v, want %v", tt.name, err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("generator.Generate(): %s = %v, want %v", tt.name, got, tt.want)
			}

		})
	}

}

func TestJWTPasswordTokenGenerator_Decode(t *testing.T) {
	generator := service.NewJWTPasswordTokenGenerator([]byte("secret"))

	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "invalid token",
			args: args{
				token: "notvalid",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "valid token",
			args: args{
				token: func() string {
					tok, err := generator.Generate(1)
					if err != nil {
						t.Fatal()
					}
					return tok
				}(),
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generator.Decode(tt.args.token)

			if (err != nil) != tt.wantErr {
				t.Errorf("generator.Decode(): error: %s = %v, want %v", tt.name, err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("generator.Decode(): %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}

}
