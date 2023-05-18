package dashboard

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/validators"
)

var _ datasource.DataSource = &dataSource{}

func NewDashboardDataSource() datasource.DataSource {
	return &dataSource{}
}

type dataSource struct{}

func (d dataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dashboard"
}

func (d dataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"rendered_json": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
			},
			"editable": schema.BoolAttribute{
				MarkdownDescription: "",
				Optional:            true,
			},
			"graph_tooltip": schema.Int64Attribute{
				MarkdownDescription: "",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2),
				},
			},
			"live_now": schema.BoolAttribute{
				MarkdownDescription: "Whether to continuously re-draw panels where the time range references 'now'",
				Optional:            true,
			},
			"panels": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "",
				Optional:            true,
			},
			"panel_json": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
			},
			"refresh": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Validators: []validator.String{
					validators.DurationStringValidator(),
				},
			},
			"style": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("light", "dark", "auto"),
				},
			},
			"tags": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "",
				Optional:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("browser", "utc"),
				},
			},
			"title": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"week_start": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("saturday", "sunday", "monday"),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"time": schema.SingleNestedBlock{
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"from": schema.StringAttribute{
						MarkdownDescription: "",
						Required:            true,
					},
					"to": schema.StringAttribute{
						MarkdownDescription: "",
						Required:            true,
					},
				},
			},
			"timepicker": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"hidden": schema.BoolAttribute{
							MarkdownDescription: "",
							Optional:            true,
						},
						"now_delay": schema.StringAttribute{
							MarkdownDescription: "",
							Optional:            true,
							Validators: []validator.String{
								validators.DurationStringValidator(),
							},
						},
						"time_options": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "",
							Optional:            true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(validators.DurationStringValidator()),
							},
						},
						"refresh_intervals": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "",
							Optional:            true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(validators.DurationStringValidator()),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
		},
	}
}

func (d dataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data model
	diags := request.Config.Get(ctx, &data)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Render the dashboard
	renderedJson, diags := RenderJson(ctx, data)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Set the state
	data.RenderedJson = types.StringValue(renderedJson)
	diags = response.State.Set(ctx, data)
	response.Diagnostics.Append(diags...)
}
