package panel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

type model struct {
	RenderedJson types.String             `tfsdk:"rendered_json"`
	NextPosition *utils.PanelNextPosition `tfsdk:"next_position"`

	Title      types.String          `tfsdk:"title"`
	Type       types.String          `tfsdk:"type"`
	Datasource utils.DatasourceModel `tfsdk:"datasource"`

	Targets []modelTarget `tfsdk:"targets"`

	Position *utils.PanelPosition `tfsdk:"position"`
	Size     utils.PanelSize      `tfsdk:"size"`

	ExtraJson types.String `tfsdk:"extra_json"`
}

type modelTarget struct {
	RefId types.String `tfsdk:"ref_id"`

	ExtraJson types.String `tfsdk:"extra_json"`
}
