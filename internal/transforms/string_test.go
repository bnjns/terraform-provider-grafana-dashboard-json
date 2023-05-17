package transforms

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromTerraformString(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		str := FromTerraformString(types.StringNull())

		assert.Nil(t, str)
	})

	t.Run("non-nil", func(t *testing.T) {
		str := FromTerraformString(types.StringValue("non-nil"))

		assert.Equal(t, "non-nil", *str)
	})
}
