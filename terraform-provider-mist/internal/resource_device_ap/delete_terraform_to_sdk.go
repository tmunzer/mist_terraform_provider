package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func DeleteTerraformToSdk(ctx context.Context) (mistapigo.MistDevice, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data mistapigo.SiteSetting
	data := mistapigo.NewDeviceAp()

	tmp := DeviceApResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}
	data.AdditionalProperties = unset

	mist_device := mistapigo.MistDevice{}
	mist_device.DeviceAp = data
	return mist_device, diags
}
