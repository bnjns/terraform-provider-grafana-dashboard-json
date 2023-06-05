package dashboard

import (
	"context"
	"fmt"
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
		Description: "Generates the JSON for Grafana dashboards.",
		Attributes: map[string]schema.Attribute{
			"rendered_json": schema.StringAttribute{
				MarkdownDescription: "The rendered dashboard JSON, which can be used with the `config_json` attribute of the [`grafana_dashboard` resource](https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/dashboard).",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "An optional description of the dashboard.",
				Optional:            true,
			},
			"editable": schema.BoolAttribute{
				MarkdownDescription: fmt.Sprintf("When `true`, those with Editor and Admin roles will be able to manually edit the dashboard. Defaults to `%t`.", defaultEditable),
				Optional:            true,
			},
			"graph_tooltip": schema.Int64Attribute{
				MarkdownDescription: fmt.Sprintf("Denotes how the tooltip is configured for panels in the dashboard. Use `0` for no shared crosshair or tooltip (default), `1` for shared crosshair, or `2` for shared crosshair and shared tooltip. Defaults to `%d`.", defaultGraphTooltip),
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2),
				},
			},
			"live_now": schema.BoolAttribute{
				MarkdownDescription: fmt.Sprintf("Whether to continuously re-draw panels where the time range references `now`. Defaults to `%t`.", defaultLiveNow),
				Optional:            true,
			},
			"panels": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The panels to include in the dashboard. This should be a list of JSON-encoded strings, and is designed to be used with the `grafana-dashboard-json_panel` data source. See the examples above.",
				Required:            true,
			},
			"extra_panel_json": schema.StringAttribute{
				MarkdownDescription: "Use this to add any additional panel JSON to the dashboard, for example for any configurations not supported by the `grafana-dashboard-json_panel` data source. This must be provided as a JSON-encoded string of a list of objects, and will be appended to any panels provided in the `panels` attribute. See the examples above.",
				Optional:            true,
			},
			"refresh": schema.StringAttribute{
				MarkdownDescription: "How frequently to auto-refresh panels in the dashboard. This must be a valid Go duration string (eg, `1m`) and is disabled by default.",
				Optional:            true,
				Validators: []validator.String{
					validators.DurationStringValidator(),
				},
			},
			"style": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf("The dashboard theme. Must be one of: %v. Defaults to `%s`.", validStyles, defaultStyle),
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf(validStyles...),
				},
			},
			"tags": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Any tags to associate with the dashboard.",
				Optional:    true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf("The default timezone the dashboard is displayed with. Must be one of: %v. Defaults to `%s`.", validTimezones, defaultTimezone),
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf(validTimezones...),
				},
			},
			"title": schema.StringAttribute{
				Description: "The title of the dashboard.",
				Required:    true,
			},
			"week_start": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf("The weekday to use at the start of each week. Must be one of: %v.", validWeekdays),
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf(validWeekdays...),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"time": schema.ListNestedBlock{
				Description: "The default time range to use when displaying the dashboard. A maximum of 1 block can be provided.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"from": schema.StringAttribute{
							MarkdownDescription: "A time, relative to `now`, to display any panels from. Must be in the form `now-<duration>` where `<duration>` is a valid duration string (eg, `6h`).",
							Required:            true,
						},
						"to": schema.StringAttribute{
							MarkdownDescription: "A time, relative to `now`, to display any panels until. Must be in the form `now-<duration>` where `<duration>` is an optional duration string (eg, `1m`). Usually configured to `now`.",
							Required:            true,
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"timepicker": schema.ListNestedBlock{
				Description: "Configuration for the timepicker. A maximum of 1 block can provided.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"hidden": schema.BoolAttribute{
							Description: "Whether the timepicker should be hidden.",
							Optional:    true,
						},
						"now_delay": schema.StringAttribute{
							MarkdownDescription: "Configure this to exclude recent data which may be incomplete. Must be a valid duration string.",
							Optional:            true,
							Validators: []validator.String{
								validators.DurationStringValidator(),
							},
						},
						"time_options": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "A mysterious setting that doesn't seem to do anything. Must be a list of valid duration strings.",
							Optional:            true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(validators.DurationStringValidator()),
							},
						},
						"refresh_intervals": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "A list of intervals users can configure the dashboard to auto-refresh with. Must be a list of valid duration strings.",
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
