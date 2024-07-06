package resource_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"mistapi/models"

	"github.com/google/uuid"
)

func TerraformToSdk(ctx context.Context, plan *SiteModel) (*models.Site, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.Site{}
	data.Name = plan.Name.ValueString()

	data.Address = models.ToPointer(plan.Address.ValueString())

	var data_latlng models.LatLng
	data_latlng.Lat = plan.Latlng.Lat.ValueFloat64()
	data_latlng.Lng = plan.Latlng.Lng.ValueFloat64()
	data.Latlng = models.ToPointer(data_latlng)

	data.CountryCode = plan.CountryCode.ValueStringPointer()

	data.Timezone = plan.Timezone.ValueStringPointer()

	data.Notes = plan.Notes.ValueStringPointer()

	var alarmtemplate_id uuid.UUID
	var aptemplate_id uuid.UUID
	var gatewaytemplate_id uuid.UUID
	var networktemplate_id uuid.UUID
	var rftemplate_id uuid.UUID
	var secpolicy_id uuid.UUID
	var sitetemplate_id uuid.UUID
	if !plan.AlarmtemplateId.IsNull() && !plan.AlarmtemplateId.IsUnknown() {
		alarmtemplate_id = uuid.MustParse(plan.AlarmtemplateId.ValueString())
	}
	if !plan.AptemplateId.IsNull() && !plan.AptemplateId.IsUnknown() {
		aptemplate_id = uuid.MustParse(plan.AptemplateId.ValueString())
	}
	if !plan.GatewaytemplateId.IsNull() && !plan.GatewaytemplateId.IsUnknown() {
		gatewaytemplate_id = uuid.MustParse(plan.GatewaytemplateId.ValueString())
	}
	if !plan.NetworktemplateId.IsNull() && !plan.NetworktemplateId.IsUnknown() {
		networktemplate_id = uuid.MustParse(plan.NetworktemplateId.ValueString())
	}
	if !plan.RftemplateId.IsNull() && !plan.RftemplateId.IsUnknown() {
		rftemplate_id = uuid.MustParse(plan.RftemplateId.ValueString())
	}
	if !plan.SecpolicyId.IsNull() && !plan.SecpolicyId.IsUnknown() {
		secpolicy_id = uuid.MustParse(plan.SecpolicyId.ValueString())
	}
	if !plan.SitetemplateId.IsNull() && !plan.SitetemplateId.IsUnknown() {
		sitetemplate_id = uuid.MustParse(plan.SitetemplateId.ValueString())
	}
	data.AlarmtemplateId.SetValue(&alarmtemplate_id)
	data.AptemplateId.SetValue(&aptemplate_id)
	data.GatewaytemplateId.SetValue(&gatewaytemplate_id)
	data.NetworktemplateId.SetValue(&networktemplate_id)
	data.RftemplateId.SetValue(&rftemplate_id)
	data.SecpolicyId.SetValue(&secpolicy_id)
	data.SitetemplateId.SetValue(&sitetemplate_id)

	var items []uuid.UUID
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, uuid.MustParse(item.String()))
	}
	data.SitegroupIds = items

	return &data, diags
}
