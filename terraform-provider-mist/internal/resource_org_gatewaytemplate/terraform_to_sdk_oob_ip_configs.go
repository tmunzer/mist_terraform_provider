package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func oobIpConfigNode1TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayOobIpConfigNode1 {
	tflog.Debug(ctx, "oobIpConfigsTerraformToSdk")

	data := models.GatewayOobIpConfigNode1{}

	if !d.IsNull() && !d.IsUnknown() {

		plan := NewNode1ValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Ip = plan.Ip.ValueStringPointer()
		data.Netmask = plan.Netmask.ValueStringPointer()
		data.Network = plan.Network.ValueStringPointer()

		data.Type = (*models.IpConfigTypeEnum)(plan.Node1Type.ValueStringPointer())
		data.UseMgmtVrf = plan.UseMgmtVrf.ValueBoolPointer()
		data.UseMgmtVrfForHostOut = plan.UseMgmtVrfForHostOut.ValueBoolPointer()
	}

	return &data
}
func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) *models.GatewayOobIpConfig {
	tflog.Debug(ctx, "oobIpConfigsTerraformToSdk")

	data := models.GatewayOobIpConfig{}

	data.Ip = d.Ip.ValueStringPointer()
	data.Netmask = d.Netmask.ValueStringPointer()
	data.Network = d.Network.ValueStringPointer()
	data.Node1 = oobIpConfigNode1TerraformToSdk(ctx, diags, d.Node1)
	data.Type = (*models.IpConfigTypeEnum)(d.OobIpConfigType.ValueStringPointer())
	data.UseMgmtVrf = d.UseMgmtVrf.ValueBoolPointer()
	data.UseMgmtVrfForHostOut = d.UseMgmtVrfForHostOut.ValueBoolPointer()

	return &data
}
