package resource_rftemplate

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

func band24TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band24Value) mistapigo.RftemplateRadioBand24 {

	data := mistapigo.NewRftemplateRadioBand24()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistapigo.RadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistapigo.Dot11Bandwidth24(plan.Bandwidth.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistapigo.RadioBandPreamble(plan.Preamble.ValueString()))

	return *data
}

func band5TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band5Value) mistapigo.RftemplateRadioBand5 {

	data := mistapigo.NewRftemplateRadioBand5()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistapigo.RadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistapigo.Dot11Bandwidth5(plan.Bandwidth.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistapigo.RadioBandPreamble(plan.Preamble.ValueString()))

	return *data
}

func band6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band6Value) mistapigo.RftemplateRadioBand6 {

	data := mistapigo.NewRftemplateRadioBand6()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistapigo.RadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistapigo.Dot11Bandwidth6(plan.Bandwidth.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistapigo.RadioBandPreamble(plan.Preamble.ValueString()))
	data.SetStandardPower(plan.StandardPower.ValueBool())

	return *data
}
