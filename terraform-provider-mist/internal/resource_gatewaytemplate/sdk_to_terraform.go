package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.Network) (NetworkModel, diag.Diagnostics) {
	var state NetworkModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.DisallowMistServices = types.BoolValue(data.GetDisallowMistServices())
	state.Gateway = types.StringValue(data.GetGateway())
	state.Gateway6 = types.StringValue(data.GetGateway6())
	state.InternalAccess = internalAccessSdkToTerraform(ctx, &diags, data.GetInternalAccess())
	state.InternetAccess = internetAccessSdkToTerraform(ctx, &diags, data.GetInternetAccess())
	state.Isolation = types.BoolValue(data.GetIsolation())
	state.RoutedForNetworks = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetRoutedForNetworks())
	state.Subnet = types.StringValue(data.GetSubnet())
	state.Subnet6 = types.StringValue(data.GetSubnet6())
	state.Tenants = tenantSdkToTerraform(ctx, &diags, data.GetTenants())
	state.VlanId = types.Int64Value(int64(data.GetVlanId()))
	state.VpnAccess = vpnSdkToTerraform(ctx, &diags, data.GetVpnAccess())

	return state, diags
}
