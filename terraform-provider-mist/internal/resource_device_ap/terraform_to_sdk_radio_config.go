package resource_device_ap

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func channelsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []int32 {
	var data []int32
	for _, v := range d.Elements() {
		var i interface{} = v
		d := i.(basetypes.Int64Value)
		data = append(data, int32(d.ValueInt64()))
	}
	return data
}

func band24TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ObjectValue) mistapigo.ApRadioBand24 {
	data := mistapigo.NewApRadioBand24()
	if plan.IsNull() || plan.IsUnknown() {
		return *data
	} else {
		var plan_interface interface{} = plan
		p := plan_interface.(Band24Value)
		data.SetAllowRrmDisable(p.AllowRrmDisable.ValueBool())
		data.SetAntGain(int32(p.AntGain.ValueInt64()))
		data.SetAntennaMode(mistapigo.RadioBandAntennaMode(p.AntennaMode.ValueString()))
		data.SetBandwidth(mistapigo.Dot11Bandwidth24(p.Bandwidth.ValueInt64()))
		data.SetChannels([]int32(p.Channels.String()))
		data.SetChannel(int32(p.Channel.ValueInt64()))
		data.SetDisabled(p.Disabled.ValueBool())
		data.SetPower(int32(p.Power.ValueInt64()))
		data.SetPowerMax(int32(p.PowerMax.ValueInt64()))
		data.SetPowerMin(int32(p.PowerMin.ValueInt64()))
		data.SetPreamble(mistapigo.RadioBandPreamble(p.Preamble.ValueString()))

		return *data
	}
}

func band5TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ObjectValue) mistapigo.ApRadioBand5 {
	data := mistapigo.NewApRadioBand5()
	if plan.IsNull() || plan.IsUnknown() {
		return *data
	} else {
		var plan_interface interface{} = plan
		p := plan_interface.(Band5Value)

		data.SetAllowRrmDisable(p.AllowRrmDisable.ValueBool())
		data.SetAntGain(int32(p.AntGain.ValueInt64()))
		data.SetAntennaMode(mistapigo.RadioBandAntennaMode(p.AntennaMode.ValueString()))
		data.SetBandwidth(mistapigo.Dot11Bandwidth5(p.Bandwidth.ValueInt64()))
		data.SetChannels([]int32(p.Channels.String()))
		data.SetChannel(int32(p.Channel.ValueInt64()))
		data.SetDisabled(p.Disabled.ValueBool())
		data.SetPower(int32(p.Power.ValueInt64()))
		data.SetPowerMax(int32(p.PowerMax.ValueInt64()))
		data.SetPowerMin(int32(p.PowerMin.ValueInt64()))
		data.SetPreamble(mistapigo.RadioBandPreamble(p.Preamble.ValueString()))

		return *data
	}
}

func band6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ObjectValue) mistapigo.ApRadioBand6 {
	data := mistapigo.NewApRadioBand6()
	if plan.IsNull() || plan.IsUnknown() {
		return *data
	} else {
		var plan_interface interface{} = plan
		p := plan_interface.(Band6Value)

		data.SetAllowRrmDisable(p.AllowRrmDisable.ValueBool())
		data.SetAntGain(int32(p.AntGain.ValueInt64()))
		data.SetAntennaMode(mistapigo.RadioBandAntennaMode(p.AntennaMode.ValueString()))
		data.SetBandwidth(mistapigo.Dot11Bandwidth6(p.Bandwidth.ValueInt64()))
		data.SetChannels([]int32(p.Channels.String()))
		data.SetChannel(int32(p.Channel.ValueInt64()))
		data.SetDisabled(p.Disabled.ValueBool())
		data.SetPower(int32(p.Power.ValueInt64()))
		data.SetPowerMax(int32(p.PowerMax.ValueInt64()))
		data.SetPowerMin(int32(p.PowerMin.ValueInt64()))
		data.SetPreamble(mistapigo.RadioBandPreamble(p.Preamble.ValueString()))
		data.SetStandardPower(p.StandardPower.ValueBool())

		return *data
	}
}

func radioConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan RadioConfigValue) mistapigo.ApRadio {

	data := mistapigo.NewApRadio()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain24(int32(plan.AntGain24.ValueInt64()))
	data.SetAntGain5(int32(plan.AntGain5.ValueInt64()))
	data.SetAntGain6(int32(plan.AntGain6.ValueInt64()))
	data.SetAntennaMode(mistapigo.ApRadioAntennaMode(plan.AntennaMode.ValueString()))

	band24 := band24TerraformToSdk(ctx, diags, plan.Band24)
	data.SetBand24(band24)
	data.SetBand24Usage(mistapigo.RadioBand24Usage(plan.Band24Usage.ValueString()))

	band5On24Radio := band5TerraformToSdk(ctx, diags, plan.Band5On24Radio)
	data.SetBand5On24Radio(band5On24Radio)

	band5 := band5TerraformToSdk(ctx, diags, plan.Band5)
	data.SetBand5(band5)

	band6 := band6TerraformToSdk(ctx, diags, plan.Band6)
	data.SetBand6(band6)

	data.SetIndoorUse(plan.IndoorUse.ValueBool())
	data.SetScanningEnabled(plan.ScanningEnabled.ValueBool())

	return *data
}
