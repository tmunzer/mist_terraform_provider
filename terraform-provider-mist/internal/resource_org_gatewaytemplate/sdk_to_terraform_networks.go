package resource_org_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
	mist_network "terraform-provider-mist/internal/resource_org_network"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func NetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.Network) basetypes.ListValue {
	tflog.Debug(ctx, "NetworksSdkToTerraform")
	var data_list = []NetworksValue{}

	for _, v := range d {
		data_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disallow_mist_services": types.BoolValue(v.GetDisallowMistServices()),
			"gateway":                types.StringValue(v.GetGateway()),
			"gateway6":               types.StringValue(v.GetGateway6()),
			"internal_access":        mist_network.InternalAccessSdkToTerraform(ctx, diags, v.GetInternalAccess()),
			"internet_access":        mist_network.InternetAccessSdkToTerraform(ctx, diags, v.GetInternetAccess()),
			"isolation":              types.BoolValue(v.GetIsolation()),
			"name":                   types.StringValue(v.GetName()),
			"routed_for_networks":    mist_transform.ListOfStringSdkToTerraform(ctx, v.GetRoutedForNetworks()),
			"subnet":                 types.StringValue(v.GetSubnet()),
			"subnet6":                types.StringValue(v.GetSubnet6()),
			"tenants":                mist_network.TenantSdkToTerraform(ctx, diags, v.GetTenants()),
			"vlan_id":                types.Int64Value(int64(v.GetVlanId())),
			"vpn_access":             mist_network.VpnSdkToTerraform(ctx, diags, v.GetVpnAccess()),
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
