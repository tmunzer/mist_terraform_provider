package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tunnelProviderOptionsJseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelProviderOptionsJse {
	tflog.Debug(ctx, "tunnelProviderOptionsJseTerraformToSdk")
	data := models.TunnelProviderOptionsJse{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewJseValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Name = plan.Name.ValueStringPointer()
		data.NumUsers = models.ToPointer(int(plan.NumUsers.ValueInt64()))
		return data
	}
}

func tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TunnelProviderOptionsZscalerSubLocation {
	tflog.Debug(ctx, "tunnelProviderOptionsZscalerSubLocationTerraformToSdk")
	var data_list []models.TunnelProviderOptionsZscalerSubLocation
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(SubLocationsValue)
		data := models.TunnelProviderOptionsZscalerSubLocation{}
		data.AupAcceptanceRequired = plan.AupAcceptanceRequired.ValueBoolPointer()
		data.AupExpire = models.ToPointer(int(plan.AupExpire.ValueInt64()))
		data.AupSslProxy = plan.AupSslProxy.ValueBoolPointer()
		data.DownloadMbps = models.ToPointer(int(plan.DownloadMbps.ValueInt64()))
		data.EnableAup = plan.EnableAup.ValueBoolPointer()
		data.EnableCaution = plan.EnableCaution.ValueBoolPointer()
		data.EnforceAuthentication = plan.EnforceAuthentication.ValueBoolPointer()
		data.Subnets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Subnets)
		data.UploadMbps = models.ToPointer(int(plan.UploadMbps.ValueInt64()))

		data_list = append(data_list, data)
	}
	return data_list
}

func tunnelProviderOptionsZscalerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelProviderOptionsZscaler {
	tflog.Debug(ctx, "tunnelProviderOptionsZscalerTerraformToSdk")
	data := models.TunnelProviderOptionsZscaler{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewZscalerValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.AupAcceptanceRequired = plan.AupAcceptanceRequired.ValueBoolPointer()
		data.AupExpire = models.ToPointer(int(plan.AupExpire.ValueInt64()))
		data.AupSslProxy = plan.AupSslProxy.ValueBoolPointer()
		data.DownloadMbps = models.ToPointer(int(plan.DownloadMbps.ValueInt64()))
		data.EnableAup = plan.EnableAup.ValueBoolPointer()
		data.EnableCaution = plan.EnableCaution.ValueBoolPointer()
		data.EnforceAuthentication = plan.EnforceAuthentication.ValueBoolPointer()
		data.Name = plan.Name.ValueStringPointer()

		sub_locations := tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx, diags, plan.SubLocations)
		data.SubLocations = sub_locations

		data.UploadMbps = models.ToPointer(int(plan.UploadMbps.ValueInt64()))
		data.UseXff = plan.UseXff.ValueBoolPointer()

		return data
	}
}

func tunnelProviderOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TunnelProviderOptionsValue) models.TunnelProviderOptions {
	tflog.Debug(ctx, "tunnelProviderOptionsTerraformToSdk")

	data := models.TunnelProviderOptions{}

	jse := tunnelProviderOptionsJseTerraformToSdk(ctx, diags, d.Jse)
	data.Jse = &jse

	zscaler := tunnelProviderOptionsZscalerTerraformToSdk(ctx, diags, d.Zscaler)
	data.Zscaler = &zscaler

	return data
}
