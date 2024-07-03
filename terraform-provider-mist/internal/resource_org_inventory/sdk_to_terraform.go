package resource_org_inventory

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func SdkToTerraform(ctx context.Context, orgId string, data []mistsdkgo.Inventory) (OrgInventoryModel, diag.Diagnostics) {
	var state OrgInventoryModel
	var diags diag.Diagnostics

	state.OrgId = types.StringValue(orgId)

	var devices []attr.Value
	for _, v := range data {
		dev_att := map[string]attr.Value{
			"magic":     types.StringValue(v.GetMagic()),
			"mac":       types.StringValue(v.GetMac()),
			"model":     types.StringValue(v.GetModel()),
			"org_id":    types.StringValue(v.GetOrgId()),
			"serial":    types.StringValue(v.GetSerial()),
			"site_id":   types.StringValue(v.GetSiteId()),
			"type":      types.StringValue(string(v.GetType())),
			"vc_mac":    types.StringValue(v.GetVcMac()),
			"hostname":  types.StringValue(v.GetHostname()),
			"device_id": types.StringValue(v.GetId()),
		}
		tflog.Debug(ctx, "----------")
		dev, e := NewDevicesValue(DevicesValue{}.AttributeTypes(ctx), dev_att)
		diags.Append(e...)
		devices = append(devices, dev)
	}

	devices_list, e := basetypes.NewListValue(DevicesValue{}.Type(ctx), devices)
	diags.Append(e...)
	state.Devices = devices_list

	return state, diags
}
