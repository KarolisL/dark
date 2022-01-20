package sinks

import (
	"context"
)

type Interface interface {
	Apply(ctx context.Context, filename string, body string) error
	Delete(ctx context.Context, filename string) error
}
