package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"terraform-provider-grafana-dashboard-json/internal/provider/dashboard"
	"terraform-provider-grafana-dashboard-json/internal/provider/panel"
	"terraform-provider-grafana-dashboard-json/internal/provider/row"
)

// Ensure GrafanaDashboardJsonProvider satisfies various provider interfaces.
var _ provider.Provider = &GrafanaDashboardJsonProvider{}

// GrafanaDashboardJsonProvider defines the provider implementation.
type GrafanaDashboardJsonProvider struct {
	version string
}

func (p *GrafanaDashboardJsonProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "grafana-dashboard-json"
	resp.Version = p.version
}

func (p *GrafanaDashboardJsonProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *GrafanaDashboardJsonProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *GrafanaDashboardJsonProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *GrafanaDashboardJsonProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		dashboard.NewDashboardDataSource,
		panel.NewPanelDataSource,
		row.NewRowDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GrafanaDashboardJsonProvider{
			version: version,
		}
	}
}
