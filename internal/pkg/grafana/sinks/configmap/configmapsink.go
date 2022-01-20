package configmap

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/K-Phoen/dark/internal/pkg/grafana"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type Sink = grafana.Sink

type DashboardSink struct {
	k8s           kubernetes.Interface
	namespace     string
	configMapName string
}

func NewDashboardSink(k8s kubernetes.Interface, namespace, configMapName string) *DashboardSink {
	return &DashboardSink{
		k8s:           k8s,
		namespace:     namespace,
		configMapName: configMapName,
	}
}

var _ Sink = (*DashboardSink)(nil)

func (d *DashboardSink) Delete(ctx context.Context, filename string) error {
	return d.patch(ctx, filename, "delete", "")
}

func (d *DashboardSink) Apply(ctx context.Context, filename string, dashboardBody string) error {
	return d.patch(ctx, filename, "replace", dashboardBody)
}
func (d *DashboardSink) patch(ctx context.Context, filename string, op string, value interface{}) error {
	patchOp, err := d.makePatchOp("replace", filename, value)
	if err != nil {
		return fmt.Errorf("marshalling Apply patch: %w", err)
	}
	// TODO: handle case when `data` does not exist (configmap is empty)
	if _, err := d.k8s.CoreV1().ConfigMaps(d.namespace).Patch(ctx, d.configMapName,
		types.JSONPatchType,
		patchOp,
		metav1.PatchOptions{},
	); err != nil {
		cmID := d.namespace + "/" + d.configMapName
		return fmt.Errorf("patching configmap %q file %q: %w", cmID, filename, err)
	}

	return nil
}

func (d *DashboardSink) makePatchOp(op string, filename string, value interface{}) ([]byte, error) {
	return json.Marshal([]struct {
		Op    string      `json:"op"`
		Path  string      `json:"path"`
		Value interface{} `json:"value"`
	}{
		{
			Op:    op,
			Path:  fmt.Sprintf("/data/%s", filename),
			Value: value,
		},
	})

}
