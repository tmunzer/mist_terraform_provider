package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderJseSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.TunnelProviderOptionsJse) basetypes.ObjectValue {
	r_attr_type := JseValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"name":      types.StringValue(d.GetName()),
		"num_users": types.Int64Value(int64(d.GetNumUsers())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelProviderZscalerSubLocationSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.TunnelProviderOptionsZscalerSubLocation) basetypes.ListValue {
	var data_list = []SubLocationsValue{}
	for _, v := range d {
		data_map_attr_type := SubLocationsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"aup_acceptance_required": types.BoolValue(v.GetAupAcceptanceRequired()),
			"aup_expire":              types.Int64Value(int64(v.GetAupExpire())),
			"aup_ssl_proxy":           types.BoolValue(v.GetAupSslProxy()),
			"download_mbps":           types.Int64Value(int64(v.GetDownloadMbps())),
			"enable_aup":              types.BoolValue(v.GetEnableAup()),
			"enable_caution":          types.BoolValue(v.GetEnableCaution()),
			"enforce_authentication":  types.BoolValue(v.GetEnforceAuthentication()),
			"subnets":                 mist_transform.ListOfStringSdkToTerraform(ctx, v.GetSubnets()),
			"upload_mbps":             types.Int64Value(int64(v.GetUploadMbps())),
		}
		data, e := NewSubLocationsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := SubLocationsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func tunnelProviderZscalerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.TunnelProviderOptionsZscaler) basetypes.ObjectValue {
	r_attr_type := ZscalerValue{}.AttributeTypes(ctx)

	sub_locations := tunnelProviderZscalerSubLocationSdkToTerraform(ctx, diags, d.GetSubLocations())

	r_attr_value := map[string]attr.Value{
		"aup_acceptance_required": types.BoolValue(d.GetAupAcceptanceRequired()),
		"aup_expire":              types.Int64Value(int64(d.GetAupExpire())),
		"aup_ssl_proxy":           types.BoolValue(d.GetAupSslProxy()),
		"download_mbps":           types.Int64Value(int64(d.GetDownloadMbps())),
		"enable_aup":              types.BoolValue(d.GetEnableAup()),
		"enable_caution":          types.BoolValue(d.GetEnableCaution()),
		"enforce_authentication":  types.BoolValue(d.GetEnforceAuthentication()),
		"name":                    types.StringValue(d.GetName()),
		"sub_locations":           sub_locations,
		"upload_mbps":             types.Int64Value(int64(d.GetUploadMbps())),
		"use_xff":                 types.BoolValue(d.GetUseXff()),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelProviderSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.TunnelProviderOptions) TunnelProviderOptionsValue {

	data_map_attr_type := TunnelProviderOptionsValue{}.AttributeTypes(ctx)

	jse := tunnelProviderJseSdkToTerraform(ctx, diags, d.GetJse())
	zscaler := tunnelProviderZscalerSdkToTerraform(ctx, diags, d.GetZscaler())

	data_map_value := map[string]attr.Value{
		"jse":     jse,
		"zscaler": zscaler,
	}
	data, e := NewTunnelProviderOptionsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
