package row

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

var _ datasource.DataSource = &dataSource{}

func NewRowDataSource() datasource.DataSource {
	return &dataSource{}
}

type dataSource struct{}

func (d dataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_row"
}

func (d dataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		"rendered_json": schema.StringAttribute{
			MarkdownDescription: "The JSON-encoded string of the row. This can be included in a dashboard by adding this to the `panels` attribute.",
			Computed:            true,
		},
		"title": schema.StringAttribute{
			MarkdownDescription: "The text to display.",
			Required:            true,
		},
	}
	utils.AddPositionSchema(attrs)
	utils.AddNextPositionSchema(attrs)

	response.Schema = schema.Schema{
		Description: "Generates the JSON for a row.",
		Attributes:  attrs,
	}
}

func (d dataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data model
	diags := request.Config.Get(ctx, &data)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	rendered := data.renderJson(ctx)
	response.Diagnostics.Append(rendered.Diagnostics...)
	if response.Diagnostics.HasError() {
		return
	}

	data.RenderedJson = types.StringValue(rendered.Json)
	data.NextPosition = &rendered.NextPosition
	diags = response.State.Set(ctx, data)
	response.Diagnostics.Append(diags...)
}
