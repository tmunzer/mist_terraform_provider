package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"
	mist_transform "terraform-provider-mist/internal/commons/utils"
	"terraform-provider-mist/internal/resource_org_network"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Network {
	tflog.Debug(ctx, "networksTerraformToSdk")
	var data_list []models.Network
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NetworksValue)
		data := models.Network{}

		data.DisallowMistServices = models.ToPointer(plan.DisallowMistServices.ValueBool())
		data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		data.Gateway6 = models.ToPointer(plan.Gateway6.ValueString())

		var internal_access_interface interface{} = plan.InternalAccess
		internal_access_tf := internal_access_interface.(resource_org_network.InternalAccessValue)
		data.InternalAccess = resource_org_network.InternalAccessTerraformToSdk(ctx, diags, internal_access_tf)

		var internet_access_interface interface{} = plan.InternetAccess
		internet_access_tf := internet_access_interface.(resource_org_network.InternetAccessValue)
		data.InternetAccess = resource_org_network.InternetAccessTerraformToSdk(ctx, diags, internet_access_tf)

		data.Isolation = models.ToPointer(plan.Isolation.ValueBool())
		data.Name = models.ToPointer(plan.Name.ValueString())
		data.RoutedForNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
		data.Subnet = models.ToPointer(plan.Subnet.ValueString())
		data.Subnet6 = models.ToPointer(plan.Subnet.ValueString())

		tenants := resource_org_network.TenantTerraformToSdk(ctx, diags, plan.Tenants)
		data.Tenants = tenants

		data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))

		vpn_access := resource_org_network.VpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		data.VpnAccess = vpn_access

		data_list = append(data_list, data)
	}
	return data_list
}
