package resource_site

import (
	"context"
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.Site) (SiteModel, diag.Diagnostics) {
	var state SiteModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	state.Address = types.StringValue(data.GetAddress())

	t := map[string]attr.Type{
		"lat": basetypes.NumberType{},
		"lng": basetypes.NumberType{},
	}
	v := map[string]attr.Value{
		"lat": types.NumberValue(big.NewFloat(data.Latlng.GetLat())),
		"lng": types.NumberValue(big.NewFloat(data.Latlng.GetLng())),
	}
	res, e := NewLatlngValue(t, v)
	diags.Append(e...)
	state.Latlng = res

	state.CountryCode = types.StringValue(data.GetCountryCode())
	state.Timezone = types.StringValue(data.GetTimezone())
	state.Notes = types.StringValue(data.GetNotes())

	state.AlarmtemplateId = types.StringValue(data.GetAlarmtemplateId())
	state.AptemplateId = types.StringValue(data.GetAptemplateId())
	state.GatewaytemplateId = types.StringValue(data.GetGatewaytemplateId())
	state.NetworktemplateId = types.StringValue(data.GetNetworktemplateId())
	state.RftemplateId = types.StringValue(data.GetRftemplateId())
	state.SecpolicyId = types.StringValue(data.GetSecpolicyId())
	state.SitetemplateId = types.StringValue(data.GetSitetemplateId())

	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data.GetSitegroupIds() {
		value := strings.ReplaceAll(item, "\"", "")
		items = append(items, types.StringValue(value))
	}
	state.SitegroupIds, _ = types.ListValue(items_type, items)
	return state, diags
}

func TerraformToSdk(ctx context.Context, plan *SiteModel) (mistsdkgo.Site, string, diag.Diagnostics) {
	data := *mistsdkgo.NewSite(plan.Name.ValueString())
	var diags diag.Diagnostics

	tflog.Debug(ctx, "SetAddress: "+plan.Address.ValueString())
	data.SetAddress(plan.Address.ValueString())

	var data_latlng mistsdkgo.LatLng
	lat, _ := plan.Latlng.Lat.ValueBigFloat().Float64()
	lng, _ := plan.Latlng.Lng.ValueBigFloat().Float64()
	data_latlng.SetLat(lat)
	data_latlng.SetLng(lng)
	data.SetLatlng(data_latlng)

	tflog.Debug(ctx, "SetCountryCode: "+plan.CountryCode.ValueString())
	data.SetCountryCode(plan.CountryCode.ValueString())

	tflog.Debug(ctx, "SetTimezone: "+plan.Timezone.ValueString())
	data.SetTimezone(plan.Timezone.ValueString())

	tflog.Debug(ctx, "SetNotes: "+plan.Notes.ValueString())
	data.SetNotes(plan.Notes.ValueString())

	tflog.Debug(ctx, "SetAlarmtemplateId: "+plan.AlarmtemplateId.ValueString())
	data.SetAlarmtemplateId(plan.AlarmtemplateId.ValueString())

	tflog.Debug(ctx, "SetAptemplateId: "+plan.AptemplateId.ValueString())
	data.SetAptemplateId(plan.AptemplateId.ValueString())

	tflog.Debug(ctx, "SetGatewaytemplateId: "+plan.GatewaytemplateId.ValueString())
	data.SetGatewaytemplateId(plan.GatewaytemplateId.ValueString())

	tflog.Debug(ctx, "SetNetworktemplateId: "+plan.NetworktemplateId.ValueString())
	data.SetNetworktemplateId(plan.NetworktemplateId.ValueString())

	tflog.Debug(ctx, "SetRftemplateId: "+plan.RftemplateId.ValueString())
	data.SetRftemplateId(plan.RftemplateId.ValueString())

	tflog.Debug(ctx, "SetSecpolicyId: "+plan.SecpolicyId.ValueString())
	data.SetSecpolicyId(plan.SecpolicyId.ValueString())

	tflog.Debug(ctx, "SetSitetemplateId: "+plan.SitetemplateId.ValueString())
	data.SetSitetemplateId(plan.SitetemplateId.ValueString())

	var items []string
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, item.String())
	}
	data.SetSitegroupIds(items)

	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}
