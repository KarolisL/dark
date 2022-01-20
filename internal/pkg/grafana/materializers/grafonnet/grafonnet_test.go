package grafonnet

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaterializer_fromBytes(t *testing.T) {
	type args struct {
		ctx  context.Context
		spec []byte
	}
	type test struct {
		name    string
		args    args
		want    string
		wantErr error
	}
	tests := []test{
		func() test {
			name := "sample_grafonnet"
			spec := loadFrom(fmt.Sprintf("./testing/%s.jsonnet", name))
			result := string(loadFrom(fmt.Sprintf("./testing/%s.json", name)))
			return test{
				name: name,
				args: args{
					ctx:  nil,
					spec: spec,
				},
				want: result,
			}
		}(),
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := &Materializer{}
			got, err := g.fromBytes(tt.args.ctx, tt.args.spec)

			if tt.wantErr == nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func loadFrom(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}
