package resource_org_network

import (
	"context"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.Network) (OrgNetworkModel, diag.Diagnostics) {
	var state OrgNetworkModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(*data.Name)

	state.DisallowMistServices = types.BoolValue(*data.DisallowMistServices)
	state.Gateway = types.StringValue(*data.Gateway)
	state.Gateway6 = types.StringValue(*data.Gateway6)
	state.InternalAccess = InternalAccessSdkToTerraform(ctx, &diags, *data.InternalAccess)
	state.InternetAccess = InternetAccessSdkToTerraform(ctx, &diags, *data.InternetAccess)
	state.Isolation = types.BoolValue(*data.Isolation)
	state.RoutedForNetworks = mist_transform.ListOfStringSdkToTerraform(ctx, data.RoutedForNetworks)
	state.Subnet = types.StringValue(*data.Subnet)
	state.Subnet6 = types.StringValue(*data.Subnet6)
	state.Tenants = TenantSdkToTerraform(ctx, &diags, data.Tenants)
	state.VlanId = types.Int64Value(int64(*data.VlanId))
	state.VpnAccess = VpnSdkToTerraform(ctx, &diags, data.VpnAccess)

	return state, diags
}
