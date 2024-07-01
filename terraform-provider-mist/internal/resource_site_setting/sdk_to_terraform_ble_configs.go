package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func bleConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.BleConfig) BleConfigValue {
	tflog.Debug(ctx, "bleConfigsSdkToTerraform")

	r_attr_type := BleConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"beacon_enabled":              types.BoolValue(d.GetBeaconEnabled()),
		"beacon_rate":                 types.Int64Value(int64(d.GetBeaconRate())),
		"beacon_rate_mode":            types.StringValue(string(d.GetBeaconRateMode())),
		"beam_disabled":               mist_transform.ListOfIntSdkToTerraform(ctx, d.GetBeamDisabled()),
		"custom_ble_packet_enabled":   types.BoolValue(d.GetCustomBlePacketEnabled()),
		"custom_ble_packet_frame":     types.StringValue(d.GetCustomBlePacketFrame()),
		"custom_ble_packet_freq_msec": types.Int64Value(int64(d.GetCustomBlePacketFreqMsec())),
		"eddystone_uid_adv_power":     types.Int64Value(int64(d.GetEddystoneUidAdvPower())),
		"eddystone_uid_beams":         types.StringValue(d.GetEddystoneUidBeams()),
		"eddystone_uid_enabled":       types.BoolValue(d.GetEddystoneUidEnabled()),
		"eddystone_uid_freq_msec":     types.Int64Value(int64(d.GetEddystoneUidFreqMsec())),
		"eddystone_uid_instance":      types.StringValue(d.GetEddystoneUidInstance()),
		"eddystone_uid_namespace":     types.StringValue(d.GetEddystoneUidInstance()),
		"eddystone_url_adv_power":     types.Int64Value(int64(d.GetEddystoneUidAdvPower())),
		"eddystone_url_beams":         types.StringValue(d.GetEddystoneUrlBeams()),
		"eddystone_url_enabled":       types.BoolValue(d.GetEddystoneUrlEnabled()),
		"eddystone_url_freq_msec":     types.Int64Value(int64(d.GetEddystoneUrlFreqMsec())),
		"eddystone_url_url":           types.StringValue(d.GetEddystoneUrlUrl()),
		"ibeacon_adv_power":           types.Int64Value(int64(d.GetIbeaconAdvPower())),
		"ibeacon_beams":               types.StringValue(d.GetIbeaconBeams()),
		"ibeacon_enabled":             types.BoolValue(d.GetIbeaconEnabled()),
		"ibeacon_freq_msec":           types.Int64Value(int64(d.GetIbeaconFreqMsec())),
		"ibeacon_major":               types.Int64Value(int64(d.GetIbeaconMajor())),
		"ibeacon_minor":               types.Int64Value(int64(d.GetIbeaconMinor())),
		"ibeacon_uuid":                types.StringValue(d.GetIbeaconUuid()),
		"power":                       types.Int64Value(int64(d.GetPower())),
		"power_mode":                  types.StringValue(string(d.GetPowerMode())),
	}
	r, e := NewBleConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
