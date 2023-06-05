package row

import (
	"terraform-provider-grafana-dashboard-json/internal/provider/utils"
)

type jsonModel struct {
	GridPos   utils.GridPositionJson `json:"gridPos"`
	Title     string                 `json:"title"`
	TitleSize string                 `json:"titleSize,omitempty"`
	Type      string                 `json:"type"`
}
