package audiorange

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr bool
	}{
		{
			name:  "return tokens",
			input: "bytes=0-499",
			want: []Token{
				{
					Type:    alpha,
					Literal: "bytes",
				},
				{
					Type:    equal,
					Literal: "=",
				},
				{
					Type:    number,
					Literal: "0",
				},
				{
					Type:    separator,
					Literal: "-",
				},
				{
					Type:    number,
					Literal: "499",
				},
				{
					Type: end,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)

			for _, tok := range tt.want {
				currentToken := l.NextToken()
				if !reflect.DeepEqual(tok, currentToken) {
					t.Errorf("Test failed: expected %v. received: %v", tok, currentToken)
				}
			}
		})
	}
}
