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

func positionAttr(t positionAttrType) schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "",
		Optional:            t == positionAttrTypeInput,
		Computed:            t == positionAttrTypeOutput,
		Attributes: map[string]schema.Attribute{
			"left": schema.Int64Attribute{
				MarkdownDescription: "",
				Required:            t == positionAttrTypeInput,
				Computed:            t == positionAttrTypeOutput,
				Validators: []validator.Int64{
					int64validator.Between(0, 23),
				},
			},
			"top": schema.Int64Attribute{
				MarkdownDescription: "",
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
	MarkdownDescription: "",
	Required:            true,
	Attributes: map[string]schema.Attribute{
		"height": schema.Int64Attribute{
			MarkdownDescription: "",
			Required:            true,
			Validators: []validator.Int64{
				int64validator.AtLeast(1),
			},
		},
		"width": schema.Int64Attribute{
			MarkdownDescription: "",
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 24),
			},
		},
	},
}

var nextPositionAttr = schema.SingleNestedAttribute{
	MarkdownDescription: "",
	Computed:            true,
	Attributes: map[string]schema.Attribute{
		"right":    positionAttr(positionAttrTypeOutput),
		"below":    positionAttr(positionAttrTypeOutput),
		"next_row": positionAttr(positionAttrTypeOutput),
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
	attrs[positionKey] = positionAttr(positionAttrTypeInput)
}

func AddNextPositionSchema(attrs map[string]schema.Attribute) {
	attrs[nextPositionKey] = nextPositionAttr
}
