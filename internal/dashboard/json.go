package dashboard

import (
	"terraform-provider-grafana-dashboard-json/internal/datasource"
)

type jsonModel struct {
	Annotations   *jsonAnnotations `json:"annotations,omitempty"`
	Description   *string          `json:"description,omitempty"`
	Editable      bool             `json:"editable"`
	GraphTooltip  int64            `json:"graphTooltip"`
	LiveNow       bool             `json:"liveNow"`
	Panels        []interface{}    `json:"panels"`
	Refresh       *string          `json:"refresh,omitempty"`
	SchemaVersion int64            `json:"schemaVersion"`
	Style         string           `json:"style"`
	Tags          []string         `json:"tags"`
	Time          *jsonTimeRange   `json:"time,omitempty"`
	Timepicker    *jsonTimepicker  `json:"timepicker,omitempty"`
	Timezone      string           `json:"timezone"`
	Title         string           `json:"title"`
	WeekStart     string           `json:"weekStart,omitempty"`

	// TODO: links
	// TODO: templating
}

type jsonAnnotations struct {
	List []jsonAnnotation `json:"list"`
}

type jsonAnnotation struct {
	BuiltIn   int    `json:"builtIn"`
	Enable    bool   `json:"enable"`
	Hide      bool   `json:"hide"`
	IconColor string `json:"iconColor"`
	Name      string `json:"name"`
	Type      string `json:"type"`

	Datasource datasource.JsonModel `json:"datasource"`
}

type jsonTimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type jsonTimepicker struct {
	Hidden           *bool    `json:"hidden,omitempty"`
	NowDelay         *string  `json:"nowDelay,omitempty"`
	TimeOptions      []string `json:"time_options,omitempty"`
	RefreshIntervals []string `json:"refresh_intervals,omitempty"`
}
