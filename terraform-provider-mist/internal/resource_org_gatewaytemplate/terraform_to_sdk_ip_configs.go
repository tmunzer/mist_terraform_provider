package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ipConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.GatewayTemplateIpConfig {
	tflog.Debug(ctx, "ipConfigsTerraformToSdk")
	data_map := make(map[string]mistapigo.GatewayTemplateIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpConfigsValue)

		data := *mistapigo.NewGatewayTemplateIpConfig()
		data.SetIp(plan.Ip.ValueString())
		data.SetNetmask(plan.Netmask.ValueString())
		data.SetSecondaryIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SecondaryIps))

		data_map[k] = data
	}
	return data_map
}
