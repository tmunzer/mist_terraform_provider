package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func portMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.SwitchPortMirroringProperty {
	data := make(map[string]mistapigo.SwitchPortMirroringProperty)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(PortMirroringValue)
		data_item := mistapigo.NewSwitchPortMirroringProperty()
		data_item.SetInputNetworksIngress(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputNetworksIngress))
		data_item.SetInputPortIdsEgress(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputPortIdsEgress))
		data_item.SetInputPortIdsIngress(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputPortIdsIngress))
		data_item.SetOutputNetwork(item_obj.OutputNetwork.ValueString())
		data_item.SetOutputPortId(item_obj.OutputPortId.ValueString())
		data[item_name] = *data_item
	}
	return data
}
