package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanAuth) AuthValue {

	var keys_list []attr.Value
	for _, item := range data.GetKeys() {
		value := item
		keys_list = append(keys_list, types.StringValue(*value))
	}
	keys, e := types.ListValue(basetypes.StringType{}, keys_list)
	diags.Append(e...)

	var pairwise_list []attr.Value
	for _, item := range data.GetPairwise() {
		value := string(item)
		pairwise_list = append(pairwise_list, types.StringValue(value))
	}
	pairwise, e := types.ListValue(basetypes.StringType{}, pairwise_list)
	diags.Append(e...)

	plan_attr := map[string]attr.Value{
		"anticlog_threshold":    types.Int64Value(int64(data.GetAnticlogThreshold())),
		"eap_reauth":            types.BoolValue(data.GetEapReauth()),
		"enable_mac_auth":       types.BoolValue(data.GetEnableMacAuth()),
		"key_idx":               types.Int64Value(int64(data.GetKeyIdx())),
		"keys":                  keys,
		"multi_psk_only":        types.BoolValue(data.GetMultiPskOnly()),
		"owe":                   types.StringValue(string(data.GetOwe())),
		"pairwise":              pairwise,
		"private_wlan":          types.BoolValue(data.GetPrivateWlan()),
		"psk":                   types.StringValue(data.GetPsk()),
		"type":                  types.StringValue(string(data.GetType())),
		"wep_as_secondary_auth": types.BoolValue(data.GetWepAsSecondaryAuth()),
	}
	r, e := NewAuthValue(AuthValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
