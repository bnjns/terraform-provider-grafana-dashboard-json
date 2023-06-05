package panel

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

var _ datasource.DataSource = &dataSource{}

func NewPanelDataSource() datasource.DataSource {
	return &dataSource{}
}

type dataSource struct{}

func (d dataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_panel"
}

func (d dataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		"datasource": schema.SingleNestedAttribute{
			MarkdownDescription: "",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					MarkdownDescription: "",
					Required:            true,
				},
				"uid": schema.StringAttribute{
					MarkdownDescription: "",
					Required:            true,
				},
			},
		},
		"extra_json": schema.StringAttribute{
			MarkdownDescription: "",
			Optional:            true,
		},
		"targets": schema.ListNestedAttribute{
			MarkdownDescription: "",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"ref_id": schema.StringAttribute{
						MarkdownDescription: "",
						Required:            true,
					},
					"extra_json": schema.StringAttribute{
						MarkdownDescription: "",
						Optional:            true,
					},
				},
			},
		},
		"title": schema.StringAttribute{
			MarkdownDescription: "",
			Required:            true,
		},
		"type": schema.StringAttribute{
			MarkdownDescription: "",
			Required:            true,
		},
	}
	utils.AddRenderedJsonSchema(attrs)
	utils.AddSizeAndPositionSchema(attrs)
	utils.AddNextPositionSchema(attrs)

	response.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          attrs,
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
	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}
