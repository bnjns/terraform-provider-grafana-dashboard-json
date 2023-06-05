package row

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

type model struct {
	RenderedJson types.String             `tfsdk:"rendered_json"`
	NextPosition *utils.PanelNextPosition `tfsdk:"next_position"`

	Title    types.String         `tfsdk:"title"`
	Position *utils.PanelPosition `tfsdk:"position"`
}
