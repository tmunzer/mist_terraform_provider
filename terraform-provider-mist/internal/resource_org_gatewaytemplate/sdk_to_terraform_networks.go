package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
	mist_network "terraform-provider-mist/internal/resource_org_network"
)

func NetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.Network) basetypes.ListValue {
	tflog.Debug(ctx, "NetworksSdkToTerraform")
	var data_list = []NetworksValue{}

	for _, v := range d {
		data_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disallow_mist_services": types.BoolValue(*v.DisallowMistServices),
			"gateway":                types.StringValue(*v.Gateway),
			"gateway6":               types.StringValue(*v.Gateway6),
			"internal_access":        mist_network.InternalAccessSdkToTerraform(ctx, diags, *v.InternalAccess),
			"internet_access":        mist_network.InternetAccessSdkToTerraform(ctx, diags, *v.InternetAccess),
			"isolation":              types.BoolValue(*v.Isolation),
			"name":                   types.StringValue(*v.Name),
			"routed_for_networks":    mist_transform.ListOfStringSdkToTerraform(ctx, v.RoutedForNetworks),
			"subnet":                 types.StringValue(*v.Subnet),
			"subnet6":                types.StringValue(*v.Subnet6),
			"tenants":                mist_network.TenantSdkToTerraform(ctx, diags, v.Tenants),
			"vlan_id":                types.Int64Value(int64(*v.VlanId)),
			"vpn_access":             mist_network.VpnSdkToTerraform(ctx, diags, v.VpnAccess),
		}

		data, e := NewNetworksValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := NetworksValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
