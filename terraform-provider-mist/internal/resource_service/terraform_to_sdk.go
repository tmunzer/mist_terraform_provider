package resource_service

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func actTagSpecsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.ServiceSpec {
	var data []mistsdkgo.ServiceSpec
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_state := v_interface.(SpecsValue)
		v_data := mistsdkgo.NewServiceSpec()
		v_data.SetPortRange(v_state.PortRange.ValueString())
		v_data.SetProtocol(v_state.Protocol.ValueString())
		data = append(data, *v_data)
	}
	return data
}

func TerraformToSdk(ctx context.Context, plan *ServiceModel) (mistsdkgo.Service, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewService()
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetAddresses(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Addresses))
	data.SetAppCategories(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppCategories))
	data.SetAppSubcategories(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppSubcategories))
	data.SetApps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Apps))
	data.SetDescription(plan.Description.ValueString())
	data.SetDscp(int32(plan.Dscp.ValueInt64()))
	data.SetFailoverPolicy(mistsdkgo.ServiceFailoverPolicy(plan.FailoverPolicy.ValueString()))
	data.SetHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostnames))
	data.SetMaxJitter(int32(*plan.MaxJitter.ValueInt64Pointer()))
	data.SetMaxLatency(int32(plan.MaxLatency.ValueInt64()))
	data.SetMaxLoss(int32(plan.MaxLoss.ValueInt64()))
	data.SetName(plan.Name.ValueString())
	data.SetSleEnabled(plan.SleEnabled.ValueBool())
	data.SetSpecs(actTagSpecsTerraformToSdk(ctx, &diags, plan.Specs))
	data.SetSsrRelaxedTcpStateEnforcement(plan.SsrRelaxedTcpStateEnforcement.ValueBool())
	data.SetTrafficClass(mistsdkgo.ServiceTrafficClass(plan.TrafficClass.ValueString()))
	data.SetTrafficType(plan.TrafficType.ValueString())
	data.SetUrls(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Urls))

	return data, diags
}
