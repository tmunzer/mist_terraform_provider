package resource_site

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *models.Site) (SiteModel, diag.Diagnostics) {
	var state SiteModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	var address basetypes.StringValue
	var latlng LatlngValue
	var country_code basetypes.StringValue
	var timezone basetypes.StringValue
	var notes basetypes.StringValue
	var alarmtemplate_id basetypes.StringValue
	var aptemplate_id basetypes.StringValue
	var gatewaytemplate_id basetypes.StringValue
	var networktemplate_id basetypes.StringValue
	var rftemplate_id basetypes.StringValue
	var secpolicy_id basetypes.StringValue
	var sitetemplate_id basetypes.StringValue

	if data.Address != nil {
		address = types.StringValue(*data.Address)
	}
	if data.Latlng != nil {
		t := map[string]attr.Type{
			"lat": basetypes.Float64Type{},
			"lng": basetypes.Float64Type{},
		}
		v := map[string]attr.Value{
			"lat": types.Float64Value(data.Latlng.Lat),
			"lng": types.Float64Value(data.Latlng.Lng),
		}
		res, e := NewLatlngValue(t, v)
		diags.Append(e...)
		latlng = res
	}
	if data.CountryCode != nil {
		country_code = types.StringValue(*data.CountryCode)
	}
	if data.Timezone != nil {
		timezone = types.StringValue(*data.Timezone)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}
	if !data.AlarmtemplateId.IsValueSet() {
		alarmtemplate_id = types.StringValue(data.AlarmtemplateId.Value().String())
	}
	if !data.AptemplateId.IsValueSet() {
		aptemplate_id = types.StringValue(data.AptemplateId.Value().String())
	}
	if !data.GatewaytemplateId.IsValueSet() {
		gatewaytemplate_id = types.StringValue(data.GatewaytemplateId.Value().String())
	}
	if !data.NetworktemplateId.IsValueSet() {
		networktemplate_id = types.StringValue(data.NetworktemplateId.Value().String())
	}
	if !data.RftemplateId.IsValueSet() {
		rftemplate_id = types.StringValue(data.RftemplateId.Value().String())
	}
	if !data.SecpolicyId.IsValueSet() {
		secpolicy_id = types.StringValue(data.SecpolicyId.Value().String())
	}
	if !data.SitetemplateId.IsValueSet() {
		sitetemplate_id = types.StringValue(data.SitetemplateId.Value().String())
	}
	state.Address = address
	state.Latlng = latlng
	state.CountryCode = country_code
	state.Timezone = timezone
	state.Notes = notes
	state.AlarmtemplateId = alarmtemplate_id
	state.AptemplateId = aptemplate_id
	state.GatewaytemplateId = gatewaytemplate_id
	state.NetworktemplateId = networktemplate_id
	state.RftemplateId = rftemplate_id
	state.SecpolicyId = secpolicy_id
	state.SitetemplateId = sitetemplate_id

	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data.SitegroupIds {
		items = append(items, types.StringValue(item.String()))
	}
	state.SitegroupIds, _ = types.ListValue(items_type, items)
	return state, diags
}
