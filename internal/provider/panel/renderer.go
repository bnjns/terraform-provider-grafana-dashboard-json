package panel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

func (m model) renderJson(ctx context.Context) utils.RenderedPanel {
	var diags diag.Diagnostics

	datasource := map[string]interface{}{
		"type": m.Datasource.Type.ValueString(),
		"uid":  m.Datasource.Uid.ValueString(),
	}
	gridPosition := utils.MakeGridPosition(m.Size, m.Position)
	targets, targetDiags := renderTargets(m.Targets, datasource)
	diags.Append(targetDiags...)

	panel := map[string]interface{}{
		"datasource": datasource,
		"gridPos":    gridPosition.ToJson(),
		"targets":    targets,
		"title":      m.Title.ValueString(),
		"type":       m.Type.ValueString(),
	}

	if !m.ExtraJson.IsNull() {
		var extraPanelJson map[string]interface{}
		err := json.Unmarshal([]byte(m.ExtraJson.ValueString()), &extraPanelJson)
		if err != nil {
			diags.AddError("Failed to parse extra panel JSON", err.Error())
			return utils.RenderedPanel{
				Diagnostics: diags,
			}
		}

		for k, v := range extraPanelJson {
			panel[k] = v
		}
	}

	panelJson, err := json.Marshal(panel)
	if err != nil {
		diags.AddError("Failed to serialise panel as JSON", err.Error())
		return utils.RenderedPanel{
			Diagnostics: diags,
		}
	}

	return utils.RenderedPanel{
		Json:         string(panelJson),
		NextPosition: gridPosition.CalculateNextPositions(),
		Diagnostics:  diags,
	}
}

func renderTargets(targets []modelTarget, datasource map[string]interface{}) ([]map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	renderedTargets := make([]map[string]interface{}, len(targets))

	for i, target := range targets {
		targetJson := map[string]interface{}{
			"refId":      target.RefId.ValueString(),
			"datasource": datasource,
		}

		if !target.ExtraJson.IsNull() {
			var targetExtraJson map[string]interface{}
			err := json.Unmarshal([]byte(target.ExtraJson.ValueString()), &targetExtraJson)
			if err != nil {
				diags.AddError(fmt.Sprintf("Failed to parse extra JSON for target %d", i), err.Error())
			} else {
				for k, v := range targetExtraJson {
					targetJson[k] = v
				}
			}
		}

		renderedTargets[i] = targetJson
	}

	return renderedTargets, diags
}
