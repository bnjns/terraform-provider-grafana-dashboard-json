package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDurationStringValidator(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	durationStringValidator := DurationStringValidator()

	t.Run("should have description", func(t *testing.T) {
		assert.NotEmpty(t, durationStringValidator.Description(ctx))
	})

	t.Run("should have a markdown description", func(t *testing.T) {
		assert.NotEmpty(t, durationStringValidator.Description(ctx))
	})

	t.Run("an invalid duration string should return an error", func(t *testing.T) {
		request := validator.StringRequest{
			Path:        path.Empty(),
			ConfigValue: types.StringValue("invalid"),
		}
		response := validator.StringResponse{}

		durationStringValidator.ValidateString(ctx, request, &response)

		assert.Equal(t, 1, len(response.Diagnostics))
		assert.Equal(t, "Must be a valid duration string", response.Diagnostics[0].Summary())
	})

	t.Run("a valid duration string should pass", func(t *testing.T) {
		request := validator.StringRequest{
			Path:        path.Empty(),
			ConfigValue: types.StringValue("5m"),
		}
		response := validator.StringResponse{}

		durationStringValidator.ValidateString(ctx, request, &response)

		assert.Empty(t, response.Diagnostics)
	})
}
