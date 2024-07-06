package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_list "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, mist_device *models.MistDevice) (DeviceApModel, diag.Diagnostics) {
	var state DeviceApModel
	var diags diag.Diagnostics

	data := mist_device.DeviceAp

	state.Id = types.StringValue(data.GetId())
	state.DeviceprofileId = types.StringValue(data.GetDeviceprofileId())
	state.MapId = types.StringValue(data.GetMapId())
	state.SiteId = types.StringValue(data.GetSiteId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	state.Notes = types.StringValue(data.GetNotes())

	state.Aeroscout = aeroscoutSdkToTerraform(ctx, &diags, data.GetAeroscout())

	state.BleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.GetBleConfig())

	state.Centrak = centrakSdkToTerraform(ctx, &diags, data.GetCentrak())

	state.ClientBridge = clientBridgeSdkToTerraform(ctx, &diags, data.GetClientBridge())

	state.DisableEth1 = types.BoolValue(data.GetDisableEth1())
	state.DisableEth2 = types.BoolValue(data.GetDisableEth2())
	state.DisableEth3 = types.BoolValue(data.GetDisableEth3())
	state.DisableModule = types.BoolValue(data.GetDisableModule())

	state.EslConfig = eslSdkToTerraform(ctx, &diags, data.GetEslConfig())

	state.Height = types.Float64Value(float64(data.GetHeight()))

	state.Image1Url = types.StringValue(data.GetImage1Url())
	state.Image2Url = types.StringValue(data.GetImage2Url())
	state.Image3Url = types.StringValue(data.GetImage3Url())

	state.IpConfig = ipConfigSdkToTerraform(ctx, &diags, data.GetIpConfig())

	state.Led = ledSdkToTerraform(ctx, &diags, data.GetLed())

	state.Locked = types.BoolValue(data.GetLocked())

	state.Mesh = meshSdkToTerraform(ctx, &diags, data.GetMesh())

	state.NtpServers = mist_list.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())

	state.Orientation = types.Int64Value(int64(data.GetOrientation()))

	state.PoePassthrough = types.BoolValue(data.GetPoePassthrough())

	state.PwrConfig = pwrConfigSdkToTerraform(ctx, &diags, data.GetPwrConfig())

	state.RadioConfig = radioConfigSdkToTerraform(ctx, &diags, data.GetRadioConfig())

	state.UplinkPortConfig = uplinkPortConfigSdkToTerraform(ctx, &diags, data.GetUplinkPortConfig())

	state.UsbConfig = usbConfigSdkToTerraform(ctx, &diags, data.GetUsbConfig())

	state.Vars = varsSdkToTerraform(ctx, &diags, data.GetVars())

	state.X = types.Float64Value(float64(data.GetX()))
	state.Y = types.Float64Value(float64(data.GetY()))

	return state, diags
}
