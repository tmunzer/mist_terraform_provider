package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func bleConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d BleConfigValue) *mistapigo.BleConfig {
	tflog.Debug(ctx, "bleConfigTerraformToSdk")
	data := mistapigo.NewBleConfig()

	data.SetBeaconEnabled(d.BeaconEnabled.ValueBool())
	data.SetBeaconRate(int32(d.BeaconRate.ValueInt64()))
	data.SetBeaconRateMode(mistapigo.BleConfigBeaconRateMode(d.BeaconRateMode.ValueString()))
	data.SetBeamDisabled(mist_transform.ListOfIntTerraformToSdk(ctx, d.BeamDisabled))
	data.SetCustomBlePacketEnabled(d.CustomBlePacketEnabled.ValueBool())
	data.SetCustomBlePacketFrame(d.CustomBlePacketFrame.ValueString())
	data.SetCustomBlePacketFreqMsec(int32(d.CustomBlePacketFreqMsec.ValueInt64()))
	data.SetEddystoneUidAdvPower(int32(d.EddystoneUidAdvPower.ValueInt64()))
	data.SetEddystoneUidBeams(d.EddystoneUidBeams.ValueString())
	data.SetEddystoneUidEnabled(d.EddystoneUidEnabled.ValueBool())
	data.SetEddystoneUidFreqMsec(int32(d.EddystoneUidFreqMsec.ValueInt64()))
	data.SetEddystoneUidInstance(d.EddystoneUidInstance.ValueString())
	data.SetEddystoneUidNamespace(d.EddystoneUidNamespace.ValueString())
	data.SetEddystoneUrlAdvPower(int32(d.EddystoneUrlAdvPower.ValueInt64()))
	data.SetEddystoneUrlBeams(d.EddystoneUrlBeams.ValueString())
	data.SetEddystoneUrlEnabled(d.EddystoneUrlEnabled.ValueBool())
	data.SetEddystoneUrlFreqMsec(int32(d.EddystoneUrlFreqMsec.ValueInt64()))
	data.SetEddystoneUrlUrl(d.EddystoneUrlUrl.ValueString())
	data.SetIbeaconAdvPower(int32(d.IbeaconAdvPower.ValueInt64()))
	data.SetIbeaconBeams(d.IbeaconBeams.ValueString())
	data.SetIbeaconEnabled(d.IbeaconEnabled.ValueBool())
	data.SetIbeaconFreqMsec(int32(d.IbeaconFreqMsec.ValueInt64()))
	data.SetIbeaconMajor(int32(d.IbeaconMajor.ValueInt64()))
	data.SetIbeaconMinor(int32(d.IbeaconMinor.ValueInt64()))
	data.SetIbeaconUuid(d.IbeaconUuid.ValueString())
	data.SetPower(int32(d.Power.ValueInt64()))
	data.SetPowerMode(mistapigo.BleConfigPowerMode(d.PowerMode.ValueString()))

	return data
}
