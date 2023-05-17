package dashboard

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type model struct {
	RenderedJson types.String `tfsdk:"rendered_json"`

	Description  types.String `tfsdk:"description"`
	Editable     types.Bool   `tfsdk:"editable"`
	GraphTooltip types.Int64  `tfsdk:"graph_tooltip"`
	LiveNow      types.Bool   `tfsdk:"live_now"`
	Refresh      types.String `tfsdk:"refresh"`
	Style        types.String `tfsdk:"style"`
	Tags         types.List   `tfsdk:"tags"`
	Time         timeRange    `tfsdk:"time"`
	Timepicker   []timepicker `tfsdk:"timepicker"`
	Timezone     types.String `tfsdk:"timezone"`
	Title        types.String `tfsdk:"title"`
	WeekStart    types.String `tfsdk:"week_start"`
}

type timeRange struct {
	From types.String `tfsdk:"from"`
	To   types.String `tfsdk:"to"`
}

type timepicker struct {
	Hidden           types.Bool   `tfsdk:"hidden"`
	NowDelay         types.String `tfsdk:"now_delay"`
	TimeOptions      types.List   `tfsdk:"time_options"`
	RefreshIntervals types.List   `tfsdk:"refresh_intervals"`
}
