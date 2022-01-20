package grafonnet

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/K-Phoen/dark/internal/pkg/grafana/materializers"
	"github.com/google/go-jsonnet"
)

type Materializer struct {
	libPath []string
}

var _ materializers.Interface = (*Materializer)(nil)

func New(libPath []string) *Materializer {
	if len(libPath) == 0 {
		libPath = []string{"./vendor"}
	}
	return &Materializer{libPath: libPath}
}

func (g *Materializer) FromSpec(ctx context.Context, folder, dashboardName string, spec []byte) (*materializers.Dashboard, error) {
	specz := struct {
		Jsonnet string `json:"jsonnet"`
	}{}
	err := json.Unmarshal(spec, &specz)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling spec: %w", err)
	}
	data, err := g.fromBytes(ctx, []byte(specz.Jsonnet))
	if err != nil {
		return nil, fmt.Errorf("evaluating jsonnet: %w", err)
	}

	return &materializers.Dashboard{
		Data:   data,
		Folder: folder,
		Name:   dashboardName,
	}, nil
}

func (g *Materializer) fromBytes(ctx context.Context, spec []byte) (string, error) {
	vm := jsonnet.MakeVM()
	vm.Importer(&jsonnet.FileImporter{
		JPaths: g.libPath,
	})
	result, err := vm.EvaluateAnonymousSnippet("<fromSpec>.jsonnet", string(spec))
	if err != nil {
		return "", fmt.Errorf("evaluating jsonnet: %w", err)
	}
	return result, nil
}
