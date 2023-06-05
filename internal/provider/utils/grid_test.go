package utils

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGridPosition(t *testing.T) {
	t.Parallel()

	size := PanelSize{
		Height: types.Int64Value(5),
		Width:  types.Int64Value(10),
	}

	t.Run("nil position defaults to 0, 0", func(t *testing.T) {
		gridPosition := MakeGridPosition(size, nil)

		assert.Equal(t, size, gridPosition.Size)
		assert.False(t, gridPosition.Position.Left.IsNull())
		assert.Equal(t, int64(0), gridPosition.Position.Left.ValueInt64())
		assert.False(t, gridPosition.Position.Top.IsNull())
		assert.Equal(t, int64(0), gridPosition.Position.Top.ValueInt64())
	})
}

func TestPanelGridPosition_Json(t *testing.T) {
	t.Parallel()

	height := int64(1)
	width := int64(24)
	left := int64(8)
	top := int64(10)

	t.Run("properties should be mapped to JSON", func(t *testing.T) {

		gridPosition := PanelGridPosition{
			Size: PanelSize{
				Height: types.Int64Value(height),
				Width:  types.Int64Value(width),
			},
			Position: PanelPosition{
				Left: types.Int64Value(left),
				Top:  types.Int64Value(top),
			},
		}

		gridPositionJson := gridPosition.ToJson()
		assert.Equal(t, height, gridPositionJson.Height)
		assert.Equal(t, width, gridPositionJson.Width)
		assert.Equal(t, left, gridPositionJson.Left)
		assert.Equal(t, top, gridPositionJson.Top)
	})

	t.Run("should be serialisable to JSON", func(t *testing.T) {
		gridPositionJson := GridPositionJson{
			Height: height,
			Width:  width,
			Left:   left,
			Top:    top,
		}

		jsonStr, err := json.Marshal(gridPositionJson)
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"h":1,"w":24,"x":8,"y":10}`), jsonStr)
	})
}

func TestPanelGridPosition_CalculateNextPositions(t *testing.T) {
	t.Parallel()

	t.Run("next positions should be calculated for normal panel", func(t *testing.T) {
		gridPosition := PanelGridPosition{
			Size: PanelSize{
				Height: types.Int64Value(10),
				Width:  types.Int64Value(8),
			},
			Position: PanelPosition{
				Left: types.Int64Value(7),
				Top:  types.Int64Value(9),
			},
		}

		nextPositions := gridPosition.CalculateNextPositions()

		assert.Equal(t, int64(15), nextPositions.Right.Left.ValueInt64())
		assert.Equal(t, int64(9), nextPositions.Right.Top.ValueInt64())

		assert.Equal(t, int64(7), nextPositions.Below.Left.ValueInt64())
		assert.Equal(t, int64(19), nextPositions.Below.Top.ValueInt64())

		assert.Equal(t, int64(0), nextPositions.NextRow.Left.ValueInt64())
		assert.Equal(t, int64(19), nextPositions.NextRow.Top.ValueInt64())
	})
}
