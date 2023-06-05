package utils

import "github.com/hashicorp/terraform-plugin-framework/types"

type DatasourceModel struct {
	Type types.String `tfsdk:"type"`
	Uid  types.String `tfsdk:"uid"`
}

type DatasourceJson struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
}
