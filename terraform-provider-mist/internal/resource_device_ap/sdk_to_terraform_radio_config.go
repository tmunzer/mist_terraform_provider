package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApRadioBand24) basetypes.ObjectValue {

	data_map_attr_type := Band24Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"channel":           types.Int64Value(int64(d.GetChannel())),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
	}

	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApRadioBand5) basetypes.ObjectValue {

	data_map_attr_type := Band5Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"channel":           types.Int64Value(int64(d.GetChannel())),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApRadioBand6) basetypes.ObjectValue {

	data_map_attr_type := Band6Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"channel":           types.Int64Value(int64(d.GetChannel())),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
		"standard_power":    types.BoolValue(d.GetStandardPower()),
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func radioConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApRadio) RadioConfigValue {
	tflog.Debug(ctx, "radioConfigSdkToTerraform")

	data_map_attr_type := RadioConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable":  types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain_24":        types.Int64Value(int64(d.GetAntGain24())),
		"ant_gain_5":         types.Int64Value(int64(d.GetAntGain5())),
		"ant_gain_6":         types.Int64Value(int64(d.GetAntGain6())),
		"antenna_mode":       types.StringValue(string(d.GetAntennaMode())),
		"band_24":            band24SdkToTerraform(ctx, diags, d.GetBand24()),
		"band_24_usage":      types.StringValue(string(d.GetBand24Usage())),
		"band_5":             band5SdkToTerraform(ctx, diags, d.GetBand5()),
		"band_5_on_24_radio": band5SdkToTerraform(ctx, diags, d.GetBand5On24Radio()),
		"band_6":             band6SdkToTerraform(ctx, diags, d.GetBand6()),
		"indoor_use":         types.BoolValue(d.GetIndoorUse()),
		"scanning_enabled":   types.BoolValue(d.GetScanningEnabled()),
	}
	data, e := NewRadioConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
