package reqcontext

import (
	"context"
	"fmt"
)

var (
	SessionID         = value{key: "session"}
	CurrentUserID     = value{key: "user"}
	ClientFingerprint = value{key: "client-fingerprint"}
	RequestID         = value{key: "request-id"}
)

type value struct {
	key string
}

func (v value) Get(ctx context.Context) (string, error) {
	return getValue(ctx, v.key)
}

func (v value) Set(ctx context.Context, s string) context.Context {
	return setValue(ctx, v.key, s)
}

func getValue(ctx context.Context, k string) (string, error) {
	if v, ok := ctx.Value(k).(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("key: %s is not present in current context", k)
}

func setValue(ctx context.Context, k string, v string) context.Context {
	return context.WithValue(ctx, k, v)
}
