package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func TerraformToSdk(ctx context.Context, plan *DeviceApModel) (mistapigo.MistDevice, diag.Diagnostics) {
	data := mistapigo.NewDeviceAp()
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data.SetDeviceprofileId(plan.DeviceprofileId.ValueString())
	data.SetMapId(plan.MapId.ValueString())
	data.SetName(plan.Name.ValueString())
	data.SetNotes(plan.Notes.ValueString())

	if !plan.Aeroscout.IsNull() && !plan.Aeroscout.IsUnknown() {
		aeroscout := aeroscoutTerraformToSdk(ctx, &diags, plan.Aeroscout)
		data.SetAeroscout(aeroscout)
	} else {
		unset["-aeroscout"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		ble_config := bleConfigTerraformToSdk(ctx, &diags, plan.BleConfig)
		data.SetBleConfig(*ble_config)
	} else {
		unset["-ble_config"] = ""
	}

	if !plan.Centrak.IsNull() && !plan.Centrak.IsUnknown() {
		centrak := centrakTerraformToSdk(ctx, &diags, plan.Centrak)
		data.SetCentrak(centrak)
	} else {
		unset["-centrak"] = ""
	}

	if !plan.ClientBridge.IsNull() && !plan.ClientBridge.IsUnknown() {
		client_bridge := clientBridgeTerraformToSdk(ctx, &diags, plan.ClientBridge)
		data.SetClientBridge(client_bridge)
	} else {
		unset["-client_bridge"] = ""
	}

	data.SetDisableEth1(plan.DisableEth1.ValueBool())
	data.SetDisableEth2(plan.DisableEth2.ValueBool())
	data.SetDisableEth3(plan.DisableEth3.ValueBool())
	data.SetDisableModule(plan.DisableModule.ValueBool())

	if !plan.EslConfig.IsNull() && !plan.EslConfig.IsUnknown() {
		esl_config := eslTerraformToSdk(ctx, &diags, plan.EslConfig)
		data.SetEslConfig(esl_config)
	} else {
		unset["-esl_config"] = ""
	}

	data.SetHeight(float32(plan.Height.ValueFloat64()))

	if !plan.IpConfig.IsNull() && !plan.IpConfig.IsUnknown() {
		ip_config := ipConfigTerraformToSdk(ctx, &diags, plan.IpConfig)
		data.SetIpConfig(ip_config)
	} else {
		unset["-ip_config"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		led := ledTerraformToSdk(ctx, &diags, plan.Led)
		data.SetLed(led)
	} else {
		unset["-led"] = ""
	}

	data.SetLocked(plan.Locked.ValueBool())

	if !plan.Mesh.IsNull() && !plan.Mesh.IsUnknown() {
		mesh := meshTerraformToSdk(ctx, &diags, plan.Mesh)
		data.SetMesh(mesh)
	} else {
		unset["-mesh"] = ""
	}

	data.SetNtpServers(mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers))

	data.SetOrientation(int32(plan.Orientation.ValueInt64()))

	data.SetPoePassthrough(plan.PoePassthrough.ValueBool())

	if !plan.PwrConfig.IsNull() && !plan.PwrConfig.IsUnknown() {
		pwr_config := pwrConfigTerraformToSdk(ctx, &diags, plan.PwrConfig)
		data.SetPwrConfig(pwr_config)
	} else {
		unset["-pwr_config"] = ""
	}

	//if !plan.RadioConfig.IsNull() && !plan.RadioConfig.IsUnknown() {
	radio_config := radioConfigTerraformToSdk(ctx, &diags, plan.RadioConfig)
	data.SetRadioConfig(radio_config)
	// } else {
	// 	unset["-radio_config"] = ""
	// }

	if !plan.UplinkPortConfig.IsNull() && !plan.UplinkPortConfig.IsUnknown() {
		uplink_port_config := uplinkPortConfigTerraformToSdk(ctx, &diags, plan.UplinkPortConfig)
		data.SetUplinkPortConfig(uplink_port_config)
	} else {
		unset["-uplink_port_config"] = ""
	}

	if !plan.UsbConfig.IsNull() && !plan.UsbConfig.IsUnknown() {
		usb_config := usbConfigTerraformToSdk(ctx, &diags, plan.UsbConfig)
		data.SetUsbConfig(usb_config)
	} else {
		unset["-usb_config"] = ""
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		vars := varsTerraformToSdk(ctx, &diags, plan.Vars)
		data.SetVars(vars)
	} else {
		unset["-vars"] = ""
	}

	data.SetX(plan.X.ValueFloat64())
	data.SetY(plan.Y.ValueFloat64())

	data.AdditionalProperties = unset

	var mist_device mistapigo.MistDevice
	mist_device.DeviceAp = data
	return mist_device, diags
}
