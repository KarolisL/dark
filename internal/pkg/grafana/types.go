package grafana

import (
	"context"
)

type Sink interface {
	Apply(ctx context.Context, filename string, body string) error
	Delete(ctx context.Context, filename string) error
}
