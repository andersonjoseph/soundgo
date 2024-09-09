package audio

import (
	"context"
)

type PlayCountHandler struct {
	repository Repository
}

func (h PlayCountHandler) Add(ctx context.Context, audio Entity) {

}
