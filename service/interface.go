package service

import (
	"context"
	"github.com/Durga-chikkala/unique-request-counter/model"
)

type UniqueRequestCounter interface {
	Get(ctx context.Context, f model.Filter) error
}
