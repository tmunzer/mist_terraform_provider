package resource_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
)

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
