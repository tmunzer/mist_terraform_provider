package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_site"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteResource{}
	_ resource.ResourceWithConfigure = &siteResource{}
)

func NewSiteResource() resource.Resource {
	return &siteResource{}
}

type siteResource struct {
	client *mistsdkgo.APIClient
}

func (r *siteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site client")
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
func (r *siteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

func (r *siteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_site.SiteResourceSchema(ctx)
}

func (r *siteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Site Create")
	var plan, state resource_site.SiteModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, diags := resource_site.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId := site.GetOrgId()
	tflog.Info(ctx, "Starting Site Create for Org "+orgId)
	data, _, err := r.client.OrgsSitesAPI.CreateOrgSite(ctx, orgId).Site(site).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site",
			"Could not create site, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site.SdkToTerraform(ctx, data)
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

func (r *siteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site.SiteModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Site Read: site_id "+state.Id.ValueString())
	data, _, err := r.client.SitesAPI.GetSiteInfo(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting site",
			"Could not get site, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site.SdkToTerraform(ctx, data)
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

func (r *siteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site.SiteModel
	tflog.Info(ctx, "Starting Site Update")

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

	site, diags := resource_site.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Site Update for Site "+state.Id.ValueString())
	data, _, err := r.client.SitesAPI.UpdateSiteInfo(ctx, state.Id.ValueString()).Site(site).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating site",
			"Could not update site, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site.SdkToTerraform(ctx, data)
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

func (r *siteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site.SiteModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Site Delete: site_id "+state.Id.ValueString())
	_, err := r.client.SitesAPI.DeleteSite(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site",
			"Could not delete site, unexpected error: "+err.Error(),
		)
		return
	}
}
