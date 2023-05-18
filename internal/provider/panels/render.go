package panels

import "github.com/hashicorp/terraform-plugin-framework/diag"

type RenderedPanel struct {
	Json         string
	NextPosition PanelNextPosition
	Diagnostics  diag.Diagnostics
}
