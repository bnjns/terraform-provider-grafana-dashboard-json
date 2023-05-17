package dashboard

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"terraform-provider-grafana-dashboard-json/internal/transforms"
)

const (
	schemaVersion       int64  = 38
	defaultEditable     bool   = true
	defaultGraphTooltip int64  = 0
	defaultLiveNow      bool   = false
	defaultTimezone     string = "browser"
	defaultStyle        string = "auto"
)

func RenderJson(ctx context.Context, data model) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	editable := defaultEditable
	if !data.Editable.IsNull() {
		editable = data.Editable.ValueBool()
	}

	graphTooltip := defaultGraphTooltip
	if !data.GraphTooltip.IsNull() {
		graphTooltip = data.GraphTooltip.ValueInt64()
	}

	liveNow := defaultLiveNow
	if !data.LiveNow.IsNull() {
		liveNow = data.LiveNow.ValueBool()
	}

	style := defaultStyle
	if !data.Style.IsNull() {
		style = data.Style.ValueString()

	}

	var tags []string
	if !data.Tags.IsNull() {
		diags = append(diags, data.Tags.ElementsAs(ctx, &tags, true)...)
	}

	timezone := defaultTimezone
	if !data.Timezone.IsNull() {
		timezone = data.Timezone.ValueString()
	}

	dashboard := &jsonModel{
		Annotations:   nil, // TODO
		Description:   transforms.FromTerraformString(data.Description),
		Editable:      editable,
		GraphTooltip:  graphTooltip,
		LiveNow:       liveNow,
		Panels:        []interface{}{},
		Refresh:       transforms.FromTerraformString(data.Refresh),
		SchemaVersion: schemaVersion,
		Style:         style,
		Tags:          tags,
		Time: &jsonTimeRange{
			From: data.Time.From.ValueString(),
			To:   data.Time.To.ValueString(),
		},
		Timepicker: renderTimepicker(ctx, data.Timepicker),
		Timezone:   timezone,
		Title:      data.Title.ValueString(),
		WeekStart:  data.WeekStart.ValueString(),
	}

	dashboardJson, err := json.Marshal(dashboard)
	if err != nil {
		diags.AddError(
			"Failed to serialise dashboard as JSON",
			err.Error(),
		)
	}

	return string(dashboardJson), diags
}

func renderTimepicker(ctx context.Context, timepicker []timepicker) *jsonTimepicker {
	if len(timepicker) == 0 {
		return nil
	}

	var timeOptions []string
	if !timepicker[0].TimeOptions.IsNull() {
		timepicker[0].TimeOptions.ElementsAs(ctx, &timeOptions, true)
	}

	var refreshIntervals []string
	if !timepicker[0].RefreshIntervals.IsNull() {
		timepicker[0].RefreshIntervals.ElementsAs(ctx, &refreshIntervals, true)
	}

	return &jsonTimepicker{
		Hidden:           transforms.FromTerraformBool(timepicker[0].Hidden),
		NowDelay:         transforms.FromTerraformString(timepicker[0].NowDelay),
		TimeOptions:      timeOptions,
		RefreshIntervals: refreshIntervals,
	}
}
