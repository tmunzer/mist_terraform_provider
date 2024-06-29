package resource_rftemplate

import (
	"context"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func channelsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []int32 {
	var data []int32
	for v := range d.Elements() {
		data = append(data, int32(v))
	}
	return data
}

func band24TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band24Value) mistsdkgo.ApRadioBand24 {

	data := mistsdkgo.NewApRadioBand24()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistsdkgo.ApRadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistsdkgo.Dot11Bandwidth24(plan.Bandwidth.ValueInt64()))
	data.SetChannel(int32(plan.Channel.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistsdkgo.ApRadioBandPreamble(plan.Preamble.ValueString()))

	return *data
}

func band5TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band5Value) mistsdkgo.ApRadioBand5 {

	data := mistsdkgo.NewApRadioBand5()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistsdkgo.ApRadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistsdkgo.Dot11Bandwidth5(plan.Bandwidth.ValueInt64()))
	data.SetChannel(int32(plan.Channel.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistsdkgo.ApRadioBandPreamble(plan.Preamble.ValueString()))

	return *data
}

func band6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band6Value) mistsdkgo.ApRadioBand6 {

	data := mistsdkgo.NewApRadioBand6()

	data.SetAllowRrmDisable(plan.AllowRrmDisable.ValueBool())
	data.SetAntGain(int32(plan.AntGain.ValueInt64()))
	data.SetAntennaMode(mistsdkgo.ApRadioBandAntennaMode(plan.AntennaMode.ValueString()))
	data.SetBandwidth(mistsdkgo.Dot11Bandwidth6(plan.Bandwidth.ValueInt64()))
	data.SetChannel(int32(plan.Channel.ValueInt64()))
	data.SetChannels(channelsTerraformToSdk(ctx, diags, plan.Channels))
	data.SetDisabled(plan.Disabled.ValueBool())
	data.SetPower(int32(plan.Power.ValueInt64()))
	data.SetPowerMax(int32(plan.PowerMax.ValueInt64()))
	data.SetPowerMin(int32(plan.PowerMin.ValueInt64()))
	data.SetPreamble(mistsdkgo.ApRadioBandPreamble(plan.Preamble.ValueString()))
	data.SetStandardPower(plan.StandardPower.ValueBool())

	return *data
}
