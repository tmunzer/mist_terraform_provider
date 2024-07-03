package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_site_wxrule"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWxRuleResource{}
	_ resource.ResourceWithConfigure = &siteWxRuleResource{}
)

func NewSiteWxRule() resource.Resource {
	return &siteWxRuleResource{}
}

type siteWxRuleResource struct {
	client *mistsdkgo.APIClient
}

func (r *siteWxRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WxRule client")
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
func (r *siteWxRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wxrule"
}

func (r *siteWxRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_site_wxrule.SiteWxruleResourceSchema(ctx)
}

func (r *siteWxRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxRule Create")
	var plan, state resource_site_wxrule.SiteWxruleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wxrule, diags := resource_site_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.SitesWxRulesAPI.CreateSiteWxRule(ctx, plan.SiteId.ValueString()).WxlanRule(*wxrule).Execute()
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating WxRule",
			"Could not create WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data)
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

func (r *siteWxRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_wxrule.SiteWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxRule Read: wxrule_id "+state.Id.ValueString())
	data, _, err := r.client.SitesWxRulesAPI.GetSiteWxRule(ctx, state.SiteId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WxRule",
			"Could not get WxRule, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data)
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

func (r *siteWxRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_wxrule.SiteWxruleModel
	tflog.Info(ctx, "Starting WxRule Update")

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

	wxrule, diags := resource_site_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxRule Update for WxRule "+state.Id.ValueString())
	data, _, err := r.client.SitesWxRulesAPI.
		UpdateSiteWxRule(ctx, state.SiteId.ValueString(), state.Id.ValueString()).
		WxlanRule(*wxrule).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WxRule",
			"Could not create WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data)
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

func (r *siteWxRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wxrule.SiteWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxRule Delete: wxrule_id "+state.Id.ValueString())
	_, err := r.client.SitesWxRulesAPI.DeleteSiteWxRule(ctx, state.SiteId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WxRule",
			"Could not delete WxRule, unexpected error: "+err.Error(),
		)
		return
	}
}
