package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

type durationStringValidator struct {
	validator.String
}

func DurationStringValidator() validator.String {
	return durationStringValidator{}
}

func (v durationStringValidator) Description(ctx context.Context) string {
	return "string must be a valid duration"
}

func (v durationStringValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v durationStringValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	var val types.String
	diags := tfsdk.ValueAs(ctx, request.ConfigValue, &val)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if val.IsUnknown() || val.IsNull() {
		return
	}

	_, err := time.ParseDuration(val.ValueString())

	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Must be a valid duration string",
			"Must be a valid duration string",
		)
	}
}
