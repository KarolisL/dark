package grafana

import (
	"context"

	"github.com/grafana-tools/sdk"
)

type DashboardBody []byte
type Folder = *sdk.Folder

type Creator interface {
	FromRawSpec(ctx context.Context, folderName string, uid string, rawJSON []byte) error
	Delete(ctx context.Context, uid string) error
}

// type Creator struct {
// 	client *sdk.Client
// }
//
// type ClientConfig struct {
// 	HttpClient  *http.Client
// 	GrafanaURL  string
// 	GrafanaAuth struct {
// 		Basic string
// 	}
// }
//
// func NewCreator(client *sdk.Client) *Creator {
// 	return &Creator{client: client}
// }
//
// func (creator *Creator) FromRawSpec(ctx context.Context, folderName string, uid string, rawJSON []byte) error {
// 	spec := make(map[string]interface{})
// 	if err := json.Unmarshal(rawJSON, &spec); err != nil {
// 		return fmt.Errorf("could not unmarshall dashboard json spec: %w", err)
// 	}
//
// 	dashboardYaml, err := yaml.Marshal(spec)
// 	if err != nil {
// 		return fmt.Errorf("could not convert dashboard spec to yaml: %w", err)
// 	}
//
// 	dashboardBuilder, err := decoder.UnmarshalYAML(bytes.NewBuffer(dashboardYaml))
// 	if err != nil {
// 		return fmt.Errorf("could not unmarshall dashboard YAML spec: %w", err)
// 	}
//
// 	dashboard.UID(uid)(&dashboardBuilder)
//
// 	return creator.upsertDashboard(ctx, folderName, dashboardBuilder)
// }
//
// func (creator *Creator) Delete(ctx context.Context, uid string) error {
// 	_, err := creator.client.DeleteDashboard(ctx, uid)
// 	return err
// }
//
// func (creator *Creator) upsertDashboard(ctx context.Context, folderName string, dashboardBody DashboardBody) error {
// 	folder, err := creator.getFolder(ctx, folderName)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := creator.doUp(ctx, folder, dashboardBody); err != nil {
// 		return fmt.Errorf("could not create dashboard: %w", err)
// 	}
//
// 	return nil
// }
//
// func (creator *Creator) getFolder(ctx context.Context, folderName string) (Folder, error) {
// 	folder, err := creator.client.GetFolderByUID(ctx, folderName)
// 	if err != nil {
// 		return nil, fmt.Errorf("retrieving folder by uid %q: %w", folderName, err)
// 	}
//
// 	return &folder, nil
// }
//
// func (creator *Creator) doUp(ctx context.Context, folder Folder, body DashboardBody) error {
// 	_, err := creator.client.SetRawDashboardWithParam(ctx, sdk.RawBoardRequest{
// 		Dashboard: body,
// 		Parameters: sdk.SetDashboardParams{
// 			FolderID:   folder.ID,
// 			Overwrite:  true,
// 			PreserveId: true,
// 		},
// 	})
//
// 	if err != nil {
// 		return fmt.Errorf("upserting dashboard: %w", err)
// 	}
//
// 	return nil
// }
