package transforms

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func FromTerraformString(str types.String) *string {
	if str.IsNull() {
		return nil
	} else {
		val := str.ValueString()
		return &val
	}
}
