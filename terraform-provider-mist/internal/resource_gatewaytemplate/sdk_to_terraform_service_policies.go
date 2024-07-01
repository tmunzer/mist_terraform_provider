package resource_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func servicePolicyAppQoESdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.ServicePolicyAppqoe) basetypes.ObjectValue {
	tflog.Debug(ctx, "servicePolicyAppQoESdkToTerraform")
	r_attr_type := AppqoeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func servicePolicyEwfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.ServicePolicyEwfRule) basetypes.ListValue {
	tflog.Debug(ctx, "servicePolicyEwfSdkToTerraform")
	var data_list = []EwfValue{}
	for _, v := range d {
		data_map_attr_type := EwfValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"alert_only":    types.BoolValue(v.GetAlertOnly()),
			"block_message": types.StringValue(v.GetBlockMessage()),
			"enabled":       types.BoolValue(v.GetEnabled()),
			"profile":       types.StringValue(string(v.GetProfile())),
		}
		data, e := NewEwfValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := EwfValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func servicePolicyIdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.IdpConfig) basetypes.ObjectValue {
	tflog.Debug(ctx, "servicePolicyIdpSdkToTerraform")
	r_attr_type := IdpValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"alert_only":    types.BoolValue(d.GetAlertOnly()),
		"enabled":       types.BoolValue(d.GetEnabled()),
		"idpprofile_id": types.StringValue(d.GetIdpprofileId()),
		"profile":       types.StringValue(string(d.GetProfile())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func servicePoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.ServicePolicy) basetypes.ListValue {
	tflog.Debug(ctx, "servicePoliciesSdkToTerraform")
	var data_list = []ServicePoliciesValue{}

	for _, v := range d {
		data_map_attr_type := ServicePoliciesValue{}.AttributeTypes(ctx)

		appqoe := servicePolicyAppQoESdkToTerraform(ctx, diags, v.GetAppqoe())
		ewf := servicePolicyEwfSdkToTerraform(ctx, diags, v.GetEwf())
		idp := servicePolicyIdpSdkToTerraform(ctx, diags, v.GetIdp())
		data_map_value := map[string]attr.Value{
			"action":           types.StringValue(string(v.GetAction())),
			"appqoe":           appqoe,
			"ewf":              ewf,
			"idp":              idp,
			"local_routing":    types.BoolValue(v.GetLocalRouting()),
			"name":             types.StringValue(v.GetName()),
			"path_preferences": types.StringValue(v.GetPathPreferences()),
			"servicepolicy_id": types.StringValue(v.GetServicepolicyId()),
			"services":         mist_transform.ListOfStringSdkToTerraform(ctx, v.GetServices()),
			"tenants":          mist_transform.ListOfStringSdkToTerraform(ctx, v.GetTenants()),
		}

		data, e := NewServicePoliciesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := ServicePoliciesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
