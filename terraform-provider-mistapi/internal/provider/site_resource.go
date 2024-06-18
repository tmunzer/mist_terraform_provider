package provider

import (
	"context"
	"fmt"
	"math/big"
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mistapi/internal/resource_site"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	var plan resource_site.SiteModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	site, orgId := processSitePlan(&plan)
	data, _, err := r.client.OrgsSitesAPI.CreateOrgSite(ctx, orgId).Site(*site).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site",
			"Could not create site, unexpected error: "+err.Error(),
		)
		return
	}

	plan = processSiteData(data)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *siteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site.SiteModel
	tflog.Info(ctx, "Starting Site Read: site_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.SitesAPI.GetSiteInfo(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting site",
			"Could not get site, unexpected error: "+err.Error(),
		)
		return
	}
	state = processSiteData(data)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_site.SiteModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//resp.Diagnostics.Append(callSiteAPI(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *siteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site.SiteModel
	tflog.Info(ctx, "Starting Site Delete: site_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SitesAPI.DeleteSite(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site",
			"Could not delete site, unexpected error: "+err.Error(),
		)
		return
	}
}

func processSiteData(data *mistsdkgo.Site) resource_site.SiteModel {
	var plan resource_site.SiteModel

	plan.Name = types.StringValue(data.GetName())
	plan.Address = types.StringValue(data.GetAddress())
	plan.Latlng.Lat = types.NumberValue(big.NewFloat(float64(data.Latlng.GetLat())))
	plan.Latlng.Lng = types.NumberValue(big.NewFloat(float64(data.Latlng.GetLng())))
	plan.CountryCode = types.StringValue(data.GetCountryCode())
	plan.Timezone = types.StringValue(data.GetTimezone())
	plan.Notes = types.StringValue(data.GetNotes())

	plan.AlarmtemplateId = types.StringValue(data.GetAlarmtemplateId())
	plan.AptemplateId = types.StringValue(data.GetAlarmtemplateId())
	plan.GatewaytemplateId = types.StringValue(data.GetGatewaytemplateId())
	plan.NetworktemplateId = types.StringValue(data.GetNetworktemplateId())
	plan.RftemplateId = types.StringValue(data.GetRftemplateId())
	plan.SecpolicyId = types.StringValue(data.GetSecpolicyId())
	plan.SitetemplateId = types.StringValue(data.GetSitetemplateId())

	var items []attr.Value
	for _, item := range data.GetSitegroupIds() {
		tmp := types.StringValue(item)
		items = append(items, tmp)
	}
	plan.SitegroupIds, _ = types.ListValue(types.StringType, items)
	return plan
}

func processSitePlan(plan *resource_site.SiteModel) (*mistsdkgo.Site, string) {
	var data *mistsdkgo.Site
	data.SetName(plan.Name.ValueString())
	data.SetAddress(plan.Address.ValueString())
	data.Latlng.Lat = float32(plan.Latlng.Lat.ValueBigFloat().Acc())
	data.Latlng.Lng = float32(plan.Latlng.Lng.ValueBigFloat().Acc())
	data.SetCountryCode(plan.CountryCode.ValueString())
	data.SetTimezone(plan.Timezone.ValueString())
	data.SetNotes(plan.Notes.ValueString())

	data.SetAlarmtemplateId(plan.AlarmtemplateId.ValueString())
	data.SetAptemplateId(plan.AptemplateId.ValueString())
	data.SetGatewaytemplateId(plan.GatewaytemplateId.ValueString())
	data.SetNetworktemplateId(plan.NetworktemplateId.ValueString())
	data.SetRftemplateId(plan.RftemplateId.ValueString())
	data.SetSecpolicyId(plan.SecpolicyId.ValueString())
	data.SetSitetemplateId(plan.SitetemplateId.ValueString())

	var orgId = plan.OrgId.String()
	return data, orgId
}
