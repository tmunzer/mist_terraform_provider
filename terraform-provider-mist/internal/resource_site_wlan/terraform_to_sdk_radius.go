package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func radiusAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.RadiusAcctServer {

	var data []mistapigo.RadiusAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AcctServersValue)
		keywrap_format, _ := mistapigo.NewRadiusKeywrapFormatFromValue(srv_plan.KeywrapFormat.ValueString())
		srv_data := mistapigo.NewRadiusAcctServer(srv_plan.Host.ValueString(), int32(srv_plan.Port.ValueInt64()), srv_plan.Secret.ValueString())
		srv_data.SetKeywrapEnabled(srv_plan.KeywrapEnabled.ValueBool())
		srv_data.SetKeywrapFormat(*keywrap_format)
		srv_data.SetKeywrapKek(srv_plan.KeywrapKek.ValueString())
		srv_data.SetKeywrapMack(srv_plan.KeywrapMack.ValueString())
		data = append(data, *srv_data)
	}
	return data
}

func radiusAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.RadiusAuthServer {

	var data []mistapigo.RadiusAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AuthServersValue)
		keywrap_format, _ := mistapigo.NewRadiusKeywrapFormatFromValue(srv_plan.KeywrapFormat.ValueString())
		srv_data := mistapigo.NewRadiusAuthServer(srv_plan.Host.ValueString(), int32(srv_plan.Port.ValueInt64()), srv_plan.Secret.ValueString())
		srv_data.SetKeywrapEnabled(srv_plan.KeywrapEnabled.ValueBool())
		srv_data.SetKeywrapFormat(*keywrap_format)
		srv_data.SetKeywrapKek(srv_plan.KeywrapKek.ValueString())
		srv_data.SetKeywrapMack(srv_plan.KeywrapMack.ValueString())
		data = append(data, *srv_data)
	}
	return data
}
