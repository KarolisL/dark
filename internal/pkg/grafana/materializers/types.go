package materializers

import (
	"context"
)

type Dashboard struct {
	Data   string
	Folder string
	Name   string
}

type Interface interface {
	FromSpec(ctx context.Context, folder, dashboardName string, spec []byte) (*Dashboard, error)
}
