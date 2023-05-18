package row

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/panels"
)

type model struct {
	RenderedJson types.String              `tfsdk:"rendered_json"`
	NextPosition *panels.PanelNextPosition `tfsdk:"next_position"`

	Title    types.String          `tfsdk:"title"`
	Position *panels.PanelPosition `tfsdk:"position"`
}
