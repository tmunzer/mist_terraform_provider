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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

	site, orgId := processSitePlan(ctx, &plan)
	tflog.Info(ctx, "Starting Site Create for Org "+orgId)
	data, _, err := r.client.OrgsSitesAPI.CreateOrgSite(ctx, orgId).Site(site).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site",
			"Could not create site, unexpected error: "+err.Error(),
		)
		return
	}

	state := processSiteData(ctx, data)

	diags = resp.State.Set(ctx, state)
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
	state = processSiteData(ctx, data)
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

	siteId := state.Id.ValueString()
	site, _ := processSitePlan(ctx, &plan)
	tflog.Info(ctx, "Starting Site Update for Site "+siteId)
	data, _, err := r.client.SitesAPI.UpdateSiteInfo(ctx, siteId).Site(site).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating site",
			"Could not update site, unexpected error: "+err.Error(),
		)
		return
	}

	state = processSiteData(ctx, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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

func processSiteData(ctx context.Context, data *mistsdkgo.Site) resource_site.SiteModel {
	var state resource_site.SiteModel

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())
	state.Address = types.StringValue(data.GetAddress())

	lat := types.NumberValue(big.NewFloat(data.Latlng.GetLat()))
	lng := types.NumberValue(big.NewFloat(data.Latlng.GetLng()))
	t := make(map[string]attr.Type)
	t["lat"] = basetypes.NumberType{}
	t["lng"] = lng.Type(ctx)
	v := make(map[string]attr.Value)
	v["lat"] = lat
	v["lng"] = lng
	state.Latlng, _ = resource_site.NewLatlngValue(t, v)

	state.CountryCode = types.StringValue(data.GetCountryCode())
	state.Timezone = types.StringValue(data.GetTimezone())
	state.Notes = types.StringValue(data.GetNotes())

	state.AlarmtemplateId = types.StringValue(data.GetAlarmtemplateId())
	state.AptemplateId = types.StringValue(data.GetAptemplateId())
	state.GatewaytemplateId = types.StringValue(data.GetGatewaytemplateId())
	state.NetworktemplateId = types.StringValue(data.GetNetworktemplateId())
	state.RftemplateId = types.StringValue(data.GetRftemplateId())
	state.SecpolicyId = types.StringValue(data.GetSecpolicyId())
	state.SitetemplateId = types.StringValue(data.GetSitetemplateId())

	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data.GetSitegroupIds() {
		items = append(items, types.StringValue(item[1:37]))
	}
	state.SitegroupIds, _ = types.ListValue(items_type, items)
	return state
}

func processSitePlan(ctx context.Context, plan *resource_site.SiteModel) (mistsdkgo.Site, string) {
	data := *mistsdkgo.NewSite(plan.Name.ValueString())

	tflog.Debug(ctx, "SetAddress: "+plan.Address.ValueString())
	data.SetAddress(plan.Address.ValueString())

	if !plan.Latlng.IsNull() && !plan.Latlng.Lat.IsNull() && !plan.Latlng.Lng.IsNull() {
		lat, _ := plan.Latlng.Lat.ValueBigFloat().Float64()
		lng, _ := plan.Latlng.Lng.ValueBigFloat().Float64()
		latlng := *mistsdkgo.NewLatLng(lat, lng)
		data.SetLatlng(latlng)
	}

	tflog.Debug(ctx, "SetCountryCode: "+plan.CountryCode.ValueString())
	data.SetCountryCode(plan.CountryCode.ValueString())

	tflog.Debug(ctx, "SetTimezone: "+plan.Timezone.ValueString())
	data.SetTimezone(plan.Timezone.ValueString())

	tflog.Debug(ctx, "SetNotes: "+plan.Notes.ValueString())
	data.SetNotes(plan.Notes.ValueString())

	tflog.Debug(ctx, "SetAlarmtemplateId: "+plan.AlarmtemplateId.ValueString())
	data.SetAlarmtemplateId(plan.AlarmtemplateId.ValueString())

	tflog.Debug(ctx, "SetAptemplateId: "+plan.AptemplateId.ValueString())
	data.SetAptemplateId(plan.AptemplateId.ValueString())

	tflog.Debug(ctx, "SetGatewaytemplateId: "+plan.GatewaytemplateId.ValueString())
	data.SetGatewaytemplateId(plan.GatewaytemplateId.ValueString())

	tflog.Debug(ctx, "SetNetworktemplateId: "+plan.NetworktemplateId.ValueString())
	data.SetNetworktemplateId(plan.NetworktemplateId.ValueString())

	tflog.Debug(ctx, "SetRftemplateId: "+plan.RftemplateId.ValueString())
	data.SetRftemplateId(plan.RftemplateId.ValueString())

	tflog.Debug(ctx, "SetSecpolicyId: "+plan.SecpolicyId.ValueString())
	data.SetSecpolicyId(plan.SecpolicyId.ValueString())

	tflog.Debug(ctx, "SetSitetemplateId: "+plan.SitetemplateId.ValueString())
	data.SetSitetemplateId(plan.SitetemplateId.ValueString())

	var items []string
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, item.String())
	}
	data.SetSitegroupIds(items)

	var orgId = plan.OrgId.ValueString()
	return data, orgId
}
