package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func dnsServerRewriteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.WlanDnsServerRewrite) DnsServerRewriteValue {

	radius_groups_values := make(map[string]attr.Value)
	for k, v := range data.GetRadiusGroups() {
		radius_groups_values[k] = types.StringValue(v)
	}
	radius_groups, e := types.MapValueFrom(ctx, types.StringType, radius_groups_values)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"enabled":       types.BoolValue(data.GetEnabled()),
		"radius_groups": radius_groups,
	}
	r, e := NewDnsServerRewriteValue(DnsServerRewriteValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
