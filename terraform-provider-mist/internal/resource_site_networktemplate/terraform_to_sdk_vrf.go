package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func vrfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrfConfigValue) mistapigo.VrfConfig {
	data := mistapigo.NewVrfConfig()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}

func vrfInstanceExtraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.VrfExtraRoute {
	data := make(map[string]mistapigo.VrfExtraRoute)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(ExtraRoutesValue)
		data_item := mistapigo.NewVrfExtraRoute()
		data_item.SetVia(item_obj.Via.ValueString())
		data[item_name] = *data_item
	}
	return data
}

func vrfInstancesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.VrfInstance {
	data := make(map[string]mistapigo.VrfInstance)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(VrfInstancesValue)
		extra_routes := vrfInstanceExtraRouteTerraformToSdk(ctx, diags, item_obj.ExtraRoutes)
		data_item := mistapigo.NewVrfInstance()
		data_item.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Networks))
		data_item.SetExtraRoutes(extra_routes)
		data[item_name] = *data_item
	}
	return data
}
