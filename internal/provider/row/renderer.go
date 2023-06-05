package row

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

const (
	_type     string = "row"
	titleSize string = "h6"
)

var size = utils.PanelSize{
	Height: types.Int64Value(1),
	Width:  types.Int64Value(24),
}

func (model model) renderJson(ctx context.Context) utils.RenderedPanel {
	var diags diag.Diagnostics

	gridPosition := utils.MakeGridPosition(size, model.Position)

	row := &jsonModel{
		Type:      _type,
		Title:     model.Title.ValueString(),
		TitleSize: titleSize,
		GridPos:   gridPosition.ToJson(),
	}

	rowJson, err := json.Marshal(row)
	if err != nil {
		diags.AddError(
			"Failed to serialise row as JSON",
			err.Error(),
		)
	}

	return utils.RenderedPanel{
		Json:         string(rowJson),
		NextPosition: gridPosition.CalculateNextPositions(),
		Diagnostics:  diags,
	}
}
