package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderOptionsJseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.TunnelProviderOptionsJse {
	data := *mistsdkgo.NewTunnelProviderOptionsJse()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(JseValue)
		data.SetName(plan.Name.ValueString())
		data.SetNumUsers(int32(plan.NumUsers.ValueInt64()))
		return data
	}
}

func tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.TunnelProviderOptionsZscalerSubLocation {
	var data_list []mistsdkgo.TunnelProviderOptionsZscalerSubLocation
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(SubLocationsValue)
		data := mistsdkgo.NewTunnelProviderOptionsZscalerSubLocation()
		data.SetAupAcceptanceRequired(plan.AupAcceptanceRequired.ValueBool())
		data.SetAupExpire(int32(plan.AupExpire.ValueInt64()))
		data.SetAupSslProxy(plan.AupSslProxy.ValueBool())
		data.SetDownloadMbps(int32(plan.DownloadMbps.ValueInt64()))
		data.SetEnableAup(plan.EnableAup.ValueBool())
		data.SetEnableCaution(plan.EnableCaution.ValueBool())
		data.SetEnforceAuthentication(plan.EnforceAuthentication.ValueBool())
		data.SetSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Subnets))
		data.SetUploadMbps(int32(plan.UploadMbps.ValueInt64()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func tunnelProviderOptionsZscalerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.TunnelProviderOptionsZscaler {
	data := *mistsdkgo.NewTunnelProviderOptionsZscaler()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(ZscalerValue)
		data.SetAupAcceptanceRequired(plan.AupAcceptanceRequired.ValueBool())
		data.SetAupExpire(int32(plan.AupExpire.ValueInt64()))
		data.SetAupSslProxy(plan.AupSslProxy.ValueBool())
		data.SetDownloadMbps(int32(plan.DownloadMbps.ValueInt64()))
		data.SetEnableAup(plan.EnableAup.ValueBool())
		data.SetEnableCaution(plan.EnableCaution.ValueBool())
		data.SetEnforceAuthentication(plan.EnforceAuthentication.ValueBool())
		data.SetName(plan.Name.ValueString())

		sub_locations := tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx, diags, plan.SubLocations)
		data.SetSubLocations(sub_locations)

		data.SetUploadMbps(int32(plan.UploadMbps.ValueInt64()))
		data.SetUseXff(plan.UseXff.ValueBool())

		return data
	}
}

func tunnelProviderOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TunnelProviderOptionsValue) mistsdkgo.TunnelProviderOptions {

	data := *mistsdkgo.NewTunnelProviderOptions()

	jse := tunnelProviderOptionsJseTerraformToSdk(ctx, diags, d.Jse)
	data.SetJse(jse)

	zscaler := tunnelProviderOptionsZscalerTerraformToSdk(ctx, diags, d.Zscaler)
	data.SetZscaler(zscaler)

	return data
}
