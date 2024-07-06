package provider

import (
	"context"
	"fmt"
	"mistapi"
	"terraform-provider-mist/internal/resource_site_wlan"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWlanResource{}
	_ resource.ResourceWithConfigure = &siteWlanResource{}
)

func NewSiteWlan() resource.Resource {
	return &siteWlanResource{}
}

type siteWlanResource struct {
	client mistapi.ClientInterface
}

func (r *siteWlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Wlan client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *siteWlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wlan"
}

func (r *siteWlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_site_wlan.SiteWlanResourceSchema(ctx)
}

func (r *siteWlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Wlan Create")
	var plan, state resource_site_wlan.SiteWlanModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wlan, diags := resource_site_wlan.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.SitesWlansAPI.CreateSiteWlan(ctx, plan.SiteId.ValueString()).Wlan(wlan
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wlan.SdkToTerraform(ctx, data.Data)
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

func (r *siteWlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_wlan.SiteWlanModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Wlan Read: wlan_id "+state.Id.ValueString())
	data, err := r.client.SitesWlansAPI.GetSiteWlan(ctx, state.SiteId.ValueString(), state.Id.ValueString()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Wlan",
			"Could not get Wlan, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_wlan.SdkToTerraform(ctx, data.Data)
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

func (r *siteWlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_wlan.SiteWlanModel
	tflog.Info(ctx, "Starting Wlan Update")

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

	wlan, diags := resource_site_wlan.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Wlan Update for Wlan "+state.Id.ValueString())
	data, err := r.client.SitesWlansAPI.
		UpdateSiteWlan(ctx, state.SiteId.ValueString(), state.Id.ValueString()).
		Wlan(wlan).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wlan.SdkToTerraform(ctx, data.Data)
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

func (r *siteWlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wlan.SiteWlanModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Wlan Delete: wlan_id "+state.Id.ValueString())
	_, err := r.client.SitesWlansAPI.DeleteSiteWlan(ctx, state.SiteId.ValueString(), state.Id.ValueString()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not delete Wlan, unexpected error: "+err.Error(),
		)
		return
	}
}
