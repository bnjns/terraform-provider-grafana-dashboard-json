package transforms

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
)

func FromTerraformInt64List(l types.List) *[]int64 {
	if l.IsNull() {
		return nil
	} else {
		var newList []int64
		for _, item := range l.Elements() {
			newItem, _ := strconv.ParseInt(item.String(), 10, 64)
			newList = append(newList, newItem)
		}

		return &newList
	}
}

func FromTerraformInt(i types.Int64) *int64 {
	if i.IsNull() {
		return nil
	} else {
		val := i.ValueInt64()
		return &val
	}
}
