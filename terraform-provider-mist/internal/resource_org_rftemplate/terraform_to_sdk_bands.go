package resource_org_rftemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func channelsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []int {
	var data []int
	for _, v := range d.Elements() {
		var i interface{} = v
		d := i.(basetypes.Int64Value)
		data = append(data, int(d.ValueInt64()))
	}
	return data
}

func band24TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band24Value) *models.RftemplateRadioBand24 {

	data := models.RftemplateRadioBand24{}

	data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
	data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
	data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
	data.Bandwidth = models.ToPointer(models.Dot11Bandwidth24Enum(plan.Bandwidth.ValueInt64()))
	data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(ctx, diags, plan.Channels)))
	data.Disabled = plan.Disabled.ValueBoolPointer()
	data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
	data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
	data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
	data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))

	return &data
}

func band5TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band5Value) *models.RftemplateRadioBand5 {

	data := models.RftemplateRadioBand5{}

	data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
	data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
	data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
	data.Bandwidth = models.ToPointer(models.Dot11Bandwidth5Enum(plan.Bandwidth.ValueInt64()))
	data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(ctx, diags, plan.Channels)))
	data.Disabled = plan.Disabled.ValueBoolPointer()
	data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
	data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
	data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
	data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))

	return &data
}

func band6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Band6Value) *models.RftemplateRadioBand6 {

	data := models.RftemplateRadioBand6{}

	data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
	data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
	data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
	data.Bandwidth = models.ToPointer(models.Dot11Bandwidth6Enum(plan.Bandwidth.ValueInt64()))
	data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(ctx, diags, plan.Channels)))
	data.Disabled = plan.Disabled.ValueBoolPointer()
	data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
	data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
	data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
	data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))
	data.StandardPower = plan.StandardPower.ValueBoolPointer()

	return &data
}
