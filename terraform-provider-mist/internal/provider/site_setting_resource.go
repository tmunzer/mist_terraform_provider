package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_site_setting"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteSettingResource{}
	_ resource.ResourceWithConfigure = &siteSettingResource{}
)

func NewSiteSettingResource() resource.Resource {
	return &siteSettingResource{}
}

type siteSettingResource struct {
	client *mistsdkgo.APIClient
}

func (r *siteSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SiteSetting client")
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
func (r *siteSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_siteSetting"
}

func (r *siteSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_site_setting.SiteSettingResourceSchema(ctx)
}

func (r *siteSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting SiteSetting Create")
	var plan, state resource_site_setting.SiteSettingModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteSetting, diags := resource_site_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting SiteSetting Create for Site "+plan.SiteId.ValueString())
	data, _, err := r.client.SitesSettingAPI.UpdateSiteSettings(ctx, plan.SiteId.ValueString()).SiteSetting(siteSetting).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating siteSetting",
			"Could not create siteSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_setting.SdkToTerraform(ctx, data)
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

func (r *siteSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_setting.SiteSettingModel
	tflog.Info(ctx, "Starting SiteSetting Read: siteSetting_id "+state.SiteId.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.SitesSettingAPI.GetSiteSetting(ctx, state.SiteId.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting siteSetting",
			"Could not get siteSetting, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_setting.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_setting.SiteSettingModel
	tflog.Info(ctx, "Starting SiteSetting Update")

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

	siteSetting, diags := resource_site_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting SiteSetting Update for SiteSetting "+state.SiteId.ValueString())
	data, _, err := r.client.SitesSettingAPI.UpdateSiteSettings(ctx, state.SiteId.ValueString()).SiteSetting(siteSetting).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating siteSetting",
			"Could not update siteSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_setting.SdkToTerraform(ctx, data)
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

func (r *siteSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_setting.SiteSettingModel
	tflog.Info(ctx, "Starting SiteSetting Delete: siteSetting_id "+state.SiteId.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := *mistsdkgo.NewSiteSetting()

	_, _, err := r.client.SitesSettingAPI.UpdateSiteSettings(ctx, state.SiteId.ValueString()).SiteSetting(data).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating siteSetting",
			"Could not delete siteSetting, unexpected error: "+err.Error(),
		)
		return
	}
}
