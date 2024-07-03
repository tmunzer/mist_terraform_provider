package resource_service

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.ServiceSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}

	for _, item := range d {
		data_map_attr_type := SpecsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": types.StringValue(string(item.GetPortRange())),
			"protocol":   types.StringValue(string(item.GetProtocol())),
		}

		data, e := NewSpecsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := SpecsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func SdkToTerraform(ctx context.Context, data *mistapigo.Service) (ServiceModel, diag.Diagnostics) {
	var state ServiceModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.Addresses = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAddresses())
	state.AppCategories = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAppCategories())
	state.AppSubcategories = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAppSubcategories())
	state.Apps = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetApps())
	state.Description = types.StringValue(data.GetDescription())
	state.Dscp = types.Int64Value(int64(data.GetDscp()))
	state.FailoverPolicy = types.StringValue(string(data.GetFailoverPolicy()))
	state.Hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetHostnames())
	state.MaxJitter = types.Int64Value(int64(data.GetMaxJitter()))
	state.MaxLatency = types.Int64Value(int64(data.GetMaxLatency()))
	state.MaxLoss = types.Int64Value(int64(data.GetMaxLoss()))
	state.Name = types.StringValue(data.GetName())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.SleEnabled = types.BoolValue(data.GetSleEnabled())
	state.Specs = serviceSpecsSdkToTerraform(ctx, &diags, data.GetSpecs())
	state.SsrRelaxedTcpStateEnforcement = types.BoolValue(data.GetSsrRelaxedTcpStateEnforcement())
	state.TrafficClass = types.StringValue(string(data.GetTrafficClass()))
	state.TrafficType = types.StringValue(data.GetTrafficType())
	state.Type = types.StringValue(string(data.GetType()))
	state.Urls = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetUrls())

	return state, diags
}
