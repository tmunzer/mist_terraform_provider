package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DnsServerRewriteValue) mistsdkgo.WlanDnsServerRewrite {

	radius_groups := make(map[string]string)
	for k, v := range plan.RadiusGroups.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		radius_groups[k] = v_plan.ValueString()
	}

	data := *mistsdkgo.NewWlanDnsServerRewrite()
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetRadiusGroups(radius_groups)

	return data
}
