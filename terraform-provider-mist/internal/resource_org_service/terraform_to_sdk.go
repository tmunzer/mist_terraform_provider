package resource_org_service

import (
	"context"

	"mistapi/models"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func actTagSpecsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServiceSpec {
	var data []models.ServiceSpec
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_state := v_interface.(SpecsValue)
		v_data := models.ServiceSpec{}

		if !v_state.PortRange.IsNull() && !v_state.PortRange.IsUnknown() {
			v_data.PortRange = v_state.PortRange.ValueStringPointer()
		}

		if !v_state.Protocol.IsNull() && !v_state.Protocol.IsUnknown() {
			v_data.Protocol = v_state.Protocol.ValueStringPointer()
		}
		data = append(data, v_data)
	}
	return data
}

func TerraformToSdk(ctx context.Context, plan *OrgServiceModel) (models.Service, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Service{}

	data.Name = plan.Name.ValueStringPointer()
	data.Specs = actTagSpecsTerraformToSdk(ctx, &diags, plan.Specs)

	if !plan.Addresses.IsNull() && !plan.Addresses.IsUnknown() {
		data.Addresses = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Addresses)
	}
	if !plan.AppCategories.IsNull() && !plan.AppCategories.IsUnknown() {
		data.AppCategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppCategories)
	}
	if !plan.AppSubcategories.IsNull() && !plan.AppSubcategories.IsUnknown() {
		data.AppSubcategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppSubcategories)
	}
	if !plan.Apps.IsNull() && !plan.Apps.IsUnknown() {
		data.Apps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Apps)
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		data.Description = plan.Description.ValueStringPointer()
	}
	if !plan.Dscp.IsNull() && !plan.Dscp.IsUnknown() {
		data.Dscp = models.ToPointer(int(plan.Dscp.ValueInt64()))
	}
	if !plan.FailoverPolicy.IsNull() && !plan.FailoverPolicy.IsUnknown() {
		data.FailoverPolicy = models.ToPointer(models.ServiceFailoverPolicyEnum(plan.FailoverPolicy.ValueString()))
	}
	if !plan.Hostnames.IsNull() && !plan.Hostnames.IsUnknown() {
		data.Hostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostnames)
	}
	if !plan.MaxJitter.IsNull() && !plan.MaxJitter.IsUnknown() {
		data.MaxJitter = models.ToPointer(int(plan.MaxJitter.ValueInt64()))
	}
	if !plan.MaxLatency.IsNull() && !plan.MaxLatency.IsUnknown() {
		data.MaxLatency = models.ToPointer(int(plan.MaxLatency.ValueInt64()))
	}
	if !plan.MaxLoss.IsNull() && !plan.MaxLoss.IsUnknown() {
		data.MaxLoss = models.ToPointer(int(plan.MaxLoss.ValueInt64()))
	}
	if !plan.SleEnabled.IsNull() && !plan.SleEnabled.IsUnknown() {
		data.SleEnabled = plan.SleEnabled.ValueBoolPointer()
	}
	if !plan.SsrRelaxedTcpStateEnforcement.IsNull() && !plan.SsrRelaxedTcpStateEnforcement.IsUnknown() {
		data.SsrRelaxedTcpStateEnforcement = plan.SsrRelaxedTcpStateEnforcement.ValueBoolPointer()
	}
	if !plan.TrafficClass.IsNull() && !plan.TrafficClass.IsUnknown() {
		data.TrafficClass = models.ToPointer(models.ServiceTrafficClassEnum(plan.TrafficClass.ValueString()))
	}
	if !plan.TrafficType.IsNull() && !plan.TrafficType.IsUnknown() {
		data.TrafficType = plan.TrafficType.ValueStringPointer()
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		data.Type = (*models.ServiceTypeEnum)(plan.Type.ValueStringPointer())
	}
	if !plan.Urls.IsNull() && !plan.Urls.IsUnknown() {
		data.Urls = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Urls)
	}

	return data, diags
}
