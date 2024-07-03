package provider

import (
	"context"
	"fmt"
	"terraform-provider-mist/internal/resource_site_networktemplate"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteNetworkTemplateResource{}
	_ resource.ResourceWithConfigure = &siteNetworkTemplateResource{}
)

func NewSiteNetworkTemplate() resource.Resource {
	return &siteNetworkTemplateResource{}
}

type siteNetworkTemplateResource struct {
	client *mistapigo.APIClient
}

func (r *siteNetworkTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NetworkTemplate client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*mistapigo.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *siteNetworkTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_networktemplate"
}

func (r *siteNetworkTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_site_networktemplate.SiteNetworktemplateResourceSchema(ctx)
}

func (r *siteNetworkTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NetworkTemplate Create")
	var plan, state resource_site_networktemplate.SiteNetworktemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networktemplate, diags := resource_site_networktemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.SitesSettingAPI.UpdateSiteSettings(ctx, plan.SiteId.ValueString()).SiteSetting(networktemplate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not create NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_networktemplate.SdkToTerraform(ctx, data)
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

func (r *siteNetworkTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_networktemplate.SiteNetworktemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NetworkTemplate Read: networktemplate_id "+state.SiteId.ValueString())
	data, _, err := r.client.SitesSettingAPI.GetSiteSetting(ctx, state.SiteId.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NetworkTemplate",
			"Could not get NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_networktemplate.SdkToTerraform(ctx, data)
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

func (r *siteNetworkTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_networktemplate.SiteNetworktemplateModel
	tflog.Info(ctx, "Starting NetworkTemplate Update")

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

	networktemplate, diags := resource_site_networktemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NetworkTemplate Update for Site "+state.SiteId.ValueString())
	data, _, err := r.client.SitesSettingAPI.
		UpdateSiteSettings(ctx, state.SiteId.ValueString()).
		SiteSetting(networktemplate).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not create NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_networktemplate.SdkToTerraform(ctx, data)
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

func (r *siteNetworkTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_networktemplate.SiteNetworktemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting NetworkTemplate Delete: site_id "+state.SiteId.ValueString())
	networktemplate, diags := resource_site_networktemplate.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(diags...)

	_, _, err := r.client.SitesSettingAPI.UpdateSiteSettings(ctx, state.SiteId.ValueString()).SiteSetting(networktemplate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not delete NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
