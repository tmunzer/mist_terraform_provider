package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func dynamicPskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanDynamicPsk) DynamicPskValue {

	plan_attr := map[string]attr.Value{
		"default_psk":     types.StringValue(data.GetDefaultPsk()),
		"default_vlan_id": types.Int64Value(int64(data.GetDefaultVlanId())),
		"enabled":         types.BoolValue(data.GetEnabled()),
		"force_lookup":    types.BoolValue(data.GetForceLookup()),
		"source":          types.StringValue(string(data.GetSource())),
	}
	r, e := NewDynamicPskValue(DynamicPskValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}