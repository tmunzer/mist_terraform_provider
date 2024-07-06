package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ipConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayTemplateIpConfig {
	tflog.Debug(ctx, "ipConfigsTerraformToSdk")
	data_map := make(map[string]models.GatewayTemplateIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpConfigsValue)

		data := models.GatewayTemplateIpConfig{}
		data.Ip = models.ToPointer(plan.Ip.ValueString())
		data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		data.SecondaryIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SecondaryIps)

		data_map[k] = data
	}
	return data_map
}
