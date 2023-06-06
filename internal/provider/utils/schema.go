package utils

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type positionAttrType int

const (
	positionKey     string = "position"
	nextPositionKey string = "next_position"
	sizeKey         string = "size"
	renderedJsonKey string = "rendered_json"

	positionAttrTypeInput positionAttrType = iota
	positionAttrTypeOutput
)

func positionAttr(t positionAttrType, desc string) schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: desc,
		Optional:            t == positionAttrTypeInput,
		Computed:            t == positionAttrTypeOutput,
		Attributes: map[string]schema.Attribute{
			"left": schema.Int64Attribute{
				MarkdownDescription: "The offset from the left of the dashboard, as a column index. Must be between 0 and 23 (inclusive).",
				Required:            t == positionAttrTypeInput,
				Computed:            t == positionAttrTypeOutput,
				Validators: []validator.Int64{
					int64validator.Between(0, 23),
				},
			},
			"top": schema.Int64Attribute{
				MarkdownDescription: "The offset from the top of the dashboard (0-indexed).",
				Required:            t == positionAttrTypeInput,
				Computed:            t == positionAttrTypeOutput,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
		},
	}
}

var sizeAttr = schema.SingleNestedAttribute{
	MarkdownDescription: "The size of the panel.",
	Required:            true,
	Attributes: map[string]schema.Attribute{
		"height": schema.Int64Attribute{
			MarkdownDescription: "The height of the panel, in \"grid height\" units (each unit is 30px).",
			Required:            true,
			Validators: []validator.Int64{
				int64validator.AtLeast(1),
			},
		},
		"width": schema.Int64Attribute{
			MarkdownDescription: "The width of the panel in columns. Must be between 1 and 24 (inclusive).",
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 24),
			},
		},
	},
}

var nextPositionAttr = schema.SingleNestedAttribute{
	MarkdownDescription: "This allows you to easily align the \"next panel\" in the dashboard relative to this panel, without needing to know it's absolute position. The desired position can be passed directly to the next panel's `position` attribute.",
	Computed:            true,
	Attributes: map[string]schema.Attribute{
		"right":    positionAttr(positionAttrTypeOutput, "The position directly to the right of this panel."),
		"below":    positionAttr(positionAttrTypeOutput, "The position directly below this panel (same offset from the left)."),
		"next_row": positionAttr(positionAttrTypeOutput, "The position at the start of the next row (below this panel, but at the left of the dashboard)."),
	},
}

func AddRenderedJsonSchema(attrs map[string]schema.Attribute) {
	attrs[renderedJsonKey] = schema.StringAttribute{
		MarkdownDescription: "",
		Computed:            true,
	}
}

func AddSizeAndPositionSchema(attrs map[string]schema.Attribute) {
	AddSizeSchema(attrs)
	AddPositionSchema(attrs)
}

func AddSizeSchema(attrs map[string]schema.Attribute) {
	attrs[sizeKey] = sizeAttr
}

func AddPositionSchema(attrs map[string]schema.Attribute) {
	attrs[positionKey] = positionAttr(positionAttrTypeInput, "The position of the top left corner of the panel.")
}

func AddNextPositionSchema(attrs map[string]schema.Attribute) {
	attrs[nextPositionKey] = nextPositionAttr
}
