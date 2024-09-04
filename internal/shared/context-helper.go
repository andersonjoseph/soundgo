package shared

import (
	"context"
	"fmt"
)

type RequestContextHandler struct{}

func (h RequestContextHandler) getValue(ctx context.Context, k string) (string, error) {
	if v, ok := ctx.Value(k).(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("key: %s is not present in request context", k)
}

func (h RequestContextHandler) setValue(ctx context.Context, k string, v string) context.Context {
	return context.WithValue(ctx, k, v)
}

func (h RequestContextHandler) SetSessionID(ctx context.Context, ID string) context.Context {
	return h.setValue(ctx, "session", ID)
}

func (h RequestContextHandler) GetSessionID(ctx context.Context) (string, error) {
	return h.getValue(ctx, "session")
}

func (h RequestContextHandler) SetUserID(ctx context.Context, ID string) context.Context {
	return h.setValue(ctx, "user", ID)
}

func (h RequestContextHandler) GetUserID(ctx context.Context) (string, error) {
	return h.getValue(ctx, "user")
}

func (h RequestContextHandler) SetHost(ctx context.Context, host string) context.Context {
	return context.WithValue(ctx, "host", host)
}

func (h RequestContextHandler) GetHost(ctx context.Context) (string, error) {
	return h.getValue(ctx, "host")
}
