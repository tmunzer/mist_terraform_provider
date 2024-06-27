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
