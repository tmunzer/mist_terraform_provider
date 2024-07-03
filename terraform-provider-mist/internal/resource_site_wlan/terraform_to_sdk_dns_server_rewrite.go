package resource_site_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DnsServerRewriteValue) mistapigo.WlanDnsServerRewrite {

	radius_groups := make(map[string]string)
	for k, v := range plan.RadiusGroups.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		radius_groups[k] = v_plan.ValueString()
	}

	data := *mistapigo.NewWlanDnsServerRewrite()
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetRadiusGroups(radius_groups)

	return data
}
