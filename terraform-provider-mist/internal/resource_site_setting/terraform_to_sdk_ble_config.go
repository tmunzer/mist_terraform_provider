package resource_site_setting

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func siteSettingBleConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d BleConfigValue) *models.BleConfig {
	tflog.Debug(ctx, "siteSettingBleConfigTerraformToSdk")
	data := models.BleConfig{}

	if !d.IsNull() && !d.IsUnknown() {
		data.BeaconEnabled = d.BeaconEnabled.ValueBoolPointer()
		data.BeaconRate = models.ToPointer(int(d.BeaconRate.ValueInt64()))
		data.BeaconRateMode = models.ToPointer(models.BleConfigBeaconRateModeEnum(d.BeaconRateMode.ValueString()))
		data.BeamDisabled = mist_transform.ListOfIntTerraformToSdk(ctx, d.BeamDisabled)
		data.CustomBlePacketEnabled = d.CustomBlePacketEnabled.ValueBoolPointer()
		data.CustomBlePacketFrame = d.CustomBlePacketFrame.ValueStringPointer()
		data.CustomBlePacketFreqMsec = models.ToPointer(int(d.CustomBlePacketFreqMsec.ValueInt64()))
		data.EddystoneUidAdvPower = models.ToPointer(int(d.EddystoneUidAdvPower.ValueInt64()))
		data.EddystoneUidBeams = d.EddystoneUidBeams.ValueStringPointer()
		data.EddystoneUidEnabled = d.EddystoneUidEnabled.ValueBoolPointer()
		data.EddystoneUidFreqMsec = models.ToPointer(int(d.EddystoneUidFreqMsec.ValueInt64()))
		data.EddystoneUidInstance = d.EddystoneUidInstance.ValueStringPointer()
		data.EddystoneUidNamespace = d.EddystoneUidNamespace.ValueStringPointer()
		data.EddystoneUrlAdvPower = models.ToPointer(int(d.EddystoneUrlAdvPower.ValueInt64()))
		data.EddystoneUrlBeams = d.EddystoneUrlBeams.ValueStringPointer()
		data.EddystoneUrlEnabled = d.EddystoneUrlEnabled.ValueBoolPointer()
		data.EddystoneUrlFreqMsec = models.ToPointer(int(d.EddystoneUrlFreqMsec.ValueInt64()))
		data.EddystoneUrlUrl = d.EddystoneUrlUrl.ValueStringPointer()
		data.IbeaconAdvPower = models.ToPointer(int(d.IbeaconAdvPower.ValueInt64()))
		data.IbeaconBeams = d.IbeaconBeams.ValueStringPointer()
		data.IbeaconEnabled = d.IbeaconEnabled.ValueBoolPointer()
		data.IbeaconFreqMsec = models.ToPointer(int(d.IbeaconFreqMsec.ValueInt64()))
		data.IbeaconMajor = models.ToPointer(int(d.IbeaconMajor.ValueInt64()))
		data.IbeaconMinor = models.ToPointer(int(d.IbeaconMinor.ValueInt64()))
		data.IbeaconUuid = models.ToPointer(uuid.MustParse(d.IbeaconUuid.ValueString()))
		data.Power = models.ToPointer(int(d.Power.ValueInt64()))
		data.PowerMode = models.ToPointer(models.BleConfigPowerModeEnum(d.PowerMode.ValueString()))
	}
	return &data
}
