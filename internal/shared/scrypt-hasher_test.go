package shared

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		p string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "hash password",
			args: args{
				p: "123",
			},
			wantErr: false,
		},
		{
			name: "hash empty string password",
			args: args{
				p: "",
			},
			wantErr: false,
		},
	}

	h := ScryptHasher{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := h.Hash(tt.args.p)

			if err != nil {
				t.Log(err)
			}

			if tt.wantErr != (err != nil) {
				t.Errorf("Test failed: err was not expected. received: %v", err)
			}

			if password == tt.args.p {
				t.Errorf("Test failed: hashed password should not be equal to password")
			}
		})
	}
}
