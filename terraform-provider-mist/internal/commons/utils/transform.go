package mist_transform

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ListOfStringTerraformToSdk(ctx context.Context, list basetypes.ListValue) []string {
	var items []string
	for _, item := range list.Elements() {
		items = append(items, strings.ReplaceAll(item.String(), "\"", ""))
	}
	return items
}

func ListOfStringSdkToTerraform(ctx context.Context, data []string) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		value := strings.ReplaceAll(item, "\"", "")
		items = append(items, types.StringValue(value))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfIntTerraformToSdk(ctx context.Context, list basetypes.ListValue) []int32 {
	var items []int32
	for _, item := range list.Elements() {
		var item_interface interface{} = item
		i := item_interface.(int64)
		items = append(items, int32(i))
	}
	return items
}

func ListOfIntSdkToTerraform(ctx context.Context, data []int32) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.Int64Type{}
	for _, item := range data {
		items = append(items, types.Int64Value(int64(item)))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}
