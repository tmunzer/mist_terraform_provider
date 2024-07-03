package provider

import (
	"context"
	"fmt"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_org_sitegroup"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgSiteGroupResource{}
	_ resource.ResourceWithConfigure = &orgSiteGroupResource{}
)

func NewOrgSiteGroupResource() resource.Resource {
	return &orgSiteGroupResource{}
}

type orgSiteGroupResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgSiteGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SiteGroup client")
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
func (r *orgSiteGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sitegroup"
}

func (r *orgSiteGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_sitegroup.OrgSitegroupResourceSchema(ctx)
}

func (r *orgSiteGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting SiteGroup Create")
	var plan, state resource_org_sitegroup.OrgSitegroupModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sitegroup, orgId, diags := resource_org_sitegroup.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsSitegroupsAPI.CreateOrgSiteGroup(ctx, orgId).Sitegroup(sitegroup).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SiteGroup",
			"Could not create SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sitegroup.SdkToTerraform(data)
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

func (r *orgSiteGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_sitegroup.OrgSitegroupModel
	tflog.Info(ctx, "Starting SiteGroup Read: sitegroup_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsSitegroupsAPI.GetOrgSiteGroup(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting SiteGroup",
			"Could not get SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_sitegroup.SdkToTerraform(data)
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

func (r *orgSiteGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_sitegroup.OrgSitegroupModel
	tflog.Info(ctx, "Starting SiteGroup Update")

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

	sitegroupId := plan.Id.ValueString()
	sitegroup, orgId, diags := resource_org_sitegroup.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting SiteGroup Update for Site "+sitegroupId)
	name := *mistsdkgo.NewNameString()
	name.SetName(sitegroup.Name)
	data, _, err := r.client.OrgsSitegroupsAPI.UpdateOrgSiteGroup(ctx, orgId, sitegroupId).NameString(name).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SiteGroup",
			"Could not create SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sitegroup.SdkToTerraform(data)
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

func (r *orgSiteGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_sitegroup.OrgSitegroupModel
	tflog.Info(ctx, "Starting SiteGroup Delete: sitegroup_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsSitegroupsAPI.DeleteOrgSiteGroup(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SiteGroup",
			"Could not delete SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
}
