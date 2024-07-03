package resource_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func TerraformToSdk(ctx context.Context, plan *SiteModel) (mistapigo.Site, diag.Diagnostics) {
	data := *mistapigo.NewSite(plan.Name.ValueString())
	var diags diag.Diagnostics

	data.SetAddress(plan.Address.ValueString())

	var data_latlng mistapigo.LatLng
	data_latlng.SetLat(plan.Latlng.Lat.ValueFloat64())
	data_latlng.SetLng(plan.Latlng.Lng.ValueFloat64())
	data.SetLatlng(data_latlng)

	data.SetCountryCode(plan.CountryCode.ValueString())

	data.SetTimezone(plan.Timezone.ValueString())

	data.SetNotes(plan.Notes.ValueString())

	data.SetAlarmtemplateId(plan.AlarmtemplateId.ValueString())

	data.SetAptemplateId(plan.AptemplateId.ValueString())

	data.SetGatewaytemplateId(plan.GatewaytemplateId.ValueString())

	data.SetNetworktemplateId(plan.NetworktemplateId.ValueString())

	data.SetRftemplateId(plan.RftemplateId.ValueString())

	data.SetSecpolicyId(plan.SecpolicyId.ValueString())

	data.SetSitetemplateId(plan.SitetemplateId.ValueString())

	data.SetOrgId(plan.OrgId.ValueString())

	var items []string
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, item.String())
	}
	data.SetSitegroupIds(items)

	return data, diags
}
