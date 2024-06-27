package resource_site

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.Site) (SiteModel, diag.Diagnostics) {
	var state SiteModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	state.Address = types.StringValue(data.GetAddress())

	t := map[string]attr.Type{
		"lat": basetypes.Float64Type{},
		"lng": basetypes.Float64Type{},
	}
	v := map[string]attr.Value{
		"lat": types.Float64Value(data.Latlng.GetLat()),
		"lng": types.Float64Value(data.Latlng.GetLng()),
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
