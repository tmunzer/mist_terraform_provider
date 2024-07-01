package resource_rftemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.RftemplateRadioBand24) Band24Value {

	data_map_attr_type := Band24Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
	}

	data, e := NewBand24Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.RftemplateRadioBand5) Band5Value {

	data_map_attr_type := Band5Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
	}
	data, e := NewBand5Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.RftemplateRadioBand6) Band6Value {

	data_map_attr_type := Band6Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": types.BoolValue(d.GetAllowRrmDisable()),
		"ant_gain":          types.Int64Value(int64(d.GetAntGain())),
		"antenna_mode":      types.StringValue(string(d.GetAntennaMode())),
		"bandwidth":         types.Int64Value(int64(d.GetBandwidth())),
		"channels":          mist_transform.ListOfIntSdkToTerraform(ctx, d.GetChannels()),
		"disabled":          types.BoolValue(d.GetDisabled()),
		"power":             types.Int64Value(int64(d.GetPower())),
		"power_max":         types.Int64Value(int64(d.GetPowerMax())),
		"power_min":         types.Int64Value(int64(d.GetPowerMin())),
		"preamble":          types.StringValue(string(d.GetPreamble())),
		"standard_power":    types.BoolValue(d.GetStandardPower()),
	}
	data, e := NewBand6Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
