package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func actTagSpecsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclTagSpec {
	var data []models.AclTagSpec
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_state := v_interface.(SpecsValue)
		v_data := models.NewAclTagSpec()
		v_data.SetPortRange(v_state.PortRange.ValueString())
		v_data.SetProtocol(v_state.Protocol.ValueString())
		data = append(data, *v_data)
	}
	return data
}

func actTagsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.AclTag {
	data := make(map[string]models.AclTag)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(AclTagsValue)
		tag_type, _ := models.NewAclTagTypeFromValue(item_obj.AclTagsType.ValueString())
		tag_specs := actTagSpecsTerraformToSdk(ctx, diags, item_obj.Specs)
		data_item := models.NewAclTag(*tag_type)
		data_item.SetGbpTag(int32(item_obj.GbpTag.ValueInt64()))
		data_item.SetMacs(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Macs))
		data_item.SetNetwork(item_obj.Network.ValueString())
		data_item.SetRadiusGroup(item_obj.RadiusGroup.ValueString())
		data_item.SetSpecs(tag_specs)
		data_item.SetSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnets))
		data[item_name] = *data_item
	}
	return data
}
