package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func vrfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrfConfigValue) mistsdkgo.JunosVrfConfig {
	data := mistsdkgo.NewJunosVrfConfig()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}

func vrfInstanceExtraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.VrfExtraRoutesValue {
	data := make(map[string]mistsdkgo.VrfExtraRoutesValue)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(ExtraRoutesValue)
		data_item := mistsdkgo.NewVrfExtraRoutesValue()
		data_item.SetVia(item_obj.Via.ValueString())
		data[item_name] = *data_item
	}
	return data
}

func vrfInstancesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.VrfInstancesConfig {
	data := make(map[string]mistsdkgo.VrfInstancesConfig)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(VrfInstancesValue)
		extra_routes := vrfInstanceExtraRouteTerraformToSdk(ctx, diags, item_obj.ExtraRoutes)
		data_item := mistsdkgo.NewVrfInstancesConfig()
		data_item.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Networks))
		data_item.SetExtraRoutes(extra_routes)
		data[item_name] = *data_item
	}
	return data
}
