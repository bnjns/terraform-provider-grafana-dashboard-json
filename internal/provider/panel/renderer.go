package panel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

func (m model) renderJson(ctx context.Context) utils.RenderedPanel {
	var diags diag.Diagnostics

	datasourceJson := utils.DatasourceJson{
		Type: m.Datasource.Type.ValueString(),
		Uid:  m.Datasource.Uid.ValueString(),
	}
	gridPosition := utils.MakeGridPosition(m.Size, m.Position)
	targets, targetDiags := renderTargets(m.Targets, datasourceJson)
	diags.Append(targetDiags...)

	// Create the initial panel object from the preset properties
	panel := map[string]interface{}{
		"datasource": datasourceJson,
		"gridPos":    gridPosition.ToJson(),
		"targets":    targets,
		"title":      m.Title.ValueString(),
		"type":       m.Type.ValueString(),
	}

	// Parse the extra panel JSON, and add it to the panel object
	if !m.ExtraJson.IsNull() {
		tflog.Debug(ctx, "Parsing extra_json and adding to panel")
		var extraPanelJson map[string]interface{}
		err := json.Unmarshal([]byte(m.ExtraJson.ValueString()), &extraPanelJson)
		if err != nil {
			diags.AddError("Failed to parse extra panel JSON", err.Error())
		} else {
			for k, v := range extraPanelJson {
				panel[k] = v
			}
		}
	}

	// Render the full panel to a JSON string
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

func renderTargets(targets []modelTarget, datasource utils.DatasourceJson) ([]map[string]interface{}, diag.Diagnostics) {
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
