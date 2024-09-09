package reqcontext

import (
	"context"
	"fmt"
)

type Handler struct{}

func (h Handler) getValue(ctx context.Context, k string) (string, error) {
	if v, ok := ctx.Value(k).(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("key: %s is not present in request context", k)
}

func (h Handler) setValue(ctx context.Context, k string, v string) context.Context {
	return context.WithValue(ctx, k, v)
}

func (h Handler) SetSessionID(ctx context.Context, ID string) context.Context {
	return h.setValue(ctx, "session", ID)
}

func (h Handler) GetSessionID(ctx context.Context) (string, error) {
	return h.getValue(ctx, "session")
}

func (h Handler) SetUserID(ctx context.Context, ID string) context.Context {
	return h.setValue(ctx, "user", ID)
}

func (h Handler) GetUserID(ctx context.Context) (string, error) {
	return h.getValue(ctx, "user")
}

func (h Handler) SetHost(ctx context.Context, host string) context.Context {
	return context.WithValue(ctx, "host", host)
}

func (h Handler) GetHost(ctx context.Context) (string, error) {
	return h.getValue(ctx, "host")
}
