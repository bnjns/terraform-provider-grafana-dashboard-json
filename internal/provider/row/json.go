package row

import (
	"terraform-provider-grafana-dashboard-json/internal/provider/panels"
)

type jsonModel struct {
	GridPos   panels.GridPositionJson `json:"gridPos"`
	Title     string                  `json:"title"`
	TitleSize string                  `json:"titleSize,omitempty"`
	Type      string                  `json:"type"`
}
