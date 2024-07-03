package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_org_nacrule"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNacRuleResource{}
	_ resource.ResourceWithConfigure = &orgNacRuleResource{}
)

func NewOrgNacRule() resource.Resource {
	return &orgNacRuleResource{}
}

type orgNacRuleResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgNacRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NacRule client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*mistsdkgo.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistsdkgo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *orgNacRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nacrule"
}

func (r *orgNacRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_nacrule.OrgNacruleResourceSchema(ctx)
}

func (r *orgNacRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NacRule Create")
	var plan, state resource_org_nacrule.OrgNacruleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nacrule, diags := resource_org_nacrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsNACRulesAPI.CreateOrgNacRule(ctx, plan.OrgId.ValueString()).NacRule(nacrule).Execute()
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating NacRule",
			"Could not create NacRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgNacRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_nacrule.OrgNacruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NacRule Read: nacrule_id "+state.Id.ValueString())
	data, _, err := r.client.OrgsNACRulesAPI.GetOrgNacRule(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NacRule",
			"Could not get NacRule, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgNacRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_nacrule.OrgNacruleModel
	tflog.Info(ctx, "Starting NacRule Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nacrule, diags := resource_org_nacrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NacRule Update for NacRule "+state.Id.ValueString())
	data, _, err := r.client.OrgsNACRulesAPI.
		UpdateOrgNacRule(ctx, state.OrgId.ValueString(), state.Id.ValueString()).
		NacRule(nacrule).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NacRule",
			"Could not create NacRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgNacRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_nacrule.OrgNacruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NacRule Delete: nacrule_id "+state.Id.ValueString())
	_, err := r.client.OrgsNACRulesAPI.DeleteOrgNacRule(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NacRule",
			"Could not delete NacRule, unexpected error: "+err.Error(),
		)
		return
	}
}
