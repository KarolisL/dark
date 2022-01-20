package grafonnet

import (
	"context"
	"fmt"

	"github.com/K-Phoen/dark/internal/pkg/grafana/materializers"
	"github.com/google/go-jsonnet"
)

type Materializer struct {
}

var _ materializers.Interface = (*Materializer)(nil)

func New() *Materializer {
	return &Materializer{}
}

func (g *Materializer) FromSpec(ctx context.Context, folder, dashboardName string, spec []byte) (*materializers.Dashboard, error) {
	panic("a")
}

func (g *Materializer) fromBytes(ctx context.Context, spec []byte) (string, error) {
	vm := jsonnet.MakeVM()
	vm.Importer(&jsonnet.FileImporter{
		JPaths: []string{"./vendor"},
	})
	json, err := vm.EvaluateAnonymousSnippet("<fromSpec>.jsonnet", string(spec))
	if err != nil {
		return "", fmt.Errorf("evaluating: %w", err)
	}
	return json, nil
}
