package panels

import "github.com/hashicorp/terraform-plugin-framework/types"

type PanelSize struct {
	Height types.Int64 `tfsdk:"height"`
	Width  types.Int64 `tfsdk:"width"`
}

type PanelPosition struct {
	Left types.Int64 `tfsdk:"left"`
	Top  types.Int64 `tfsdk:"top"`
}

type PanelGridPosition struct {
	Size     PanelSize
	Position PanelPosition
}

type GridPositionJson struct {
	Height int64 `json:"h"`
	Width  int64 `json:"w"`

	Left int64 `json:"x"`
	Top  int64 `json:"y"`
}

type PanelNextPosition struct {
	Right   PanelPosition `tfsdk:"right"`
	Below   PanelPosition `tfsdk:"below"`
	NextRow PanelPosition `tfsdk:"next_row"`
}

func MakeGridPosition(size PanelSize, position *PanelPosition) PanelGridPosition {
	if position == nil {
		return PanelGridPosition{
			Size: size,
			Position: PanelPosition{
				Left: types.Int64Value(0),
				Top:  types.Int64Value(0),
			},
		}
	} else {
		return PanelGridPosition{
			Size:     size,
			Position: *position,
		}
	}
}

func (p PanelGridPosition) ToJson() GridPositionJson {
	return GridPositionJson{
		Height: p.Size.Height.ValueInt64(),
		Width:  p.Size.Width.ValueInt64(),

		Left: p.Position.Left.ValueInt64(),
		Top:  p.Position.Top.ValueInt64(),
	}
}

func (p PanelGridPosition) CalculateNextPositions() PanelNextPosition {
	return PanelNextPosition{
		Right: PanelPosition{
			Left: types.Int64Value(p.Position.Left.ValueInt64() + p.Size.Width.ValueInt64()),
			Top:  p.Position.Top,
		},
		Below: PanelPosition{
			Left: p.Position.Left,
			Top:  types.Int64Value(p.Position.Top.ValueInt64() + p.Size.Height.ValueInt64()),
		},
		NextRow: PanelPosition{
			Left: types.Int64Value(0),
			Top:  types.Int64Value(p.Position.Top.ValueInt64() + p.Size.Height.ValueInt64()),
		},
	}
}
