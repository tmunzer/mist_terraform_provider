package resource_org_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ipConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.GatewayTemplateIpConfig {
	tflog.Debug(ctx, "ipConfigsTerraformToSdk")
	data_map := make(map[string]mistsdkgo.GatewayTemplateIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpConfigsValue)

		data := *mistsdkgo.NewGatewayTemplateIpConfig()
		data.SetIp(plan.Ip.ValueString())
		data.SetNetmask(plan.Netmask.ValueString())
		data.SetSecondaryIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SecondaryIps))

		data_map[k] = data
	}
	return data_map
}
