/*
Mist API

> Version: **2406.1.10** > > Date: **July 1, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.11
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
)

// checks if the UiSettingsTile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiSettingsTile{}

// UiSettingsTile struct for UiSettingsTile
type UiSettingsTile struct {
	ChartBand *string `json:"chartBand,omitempty"`
	ChartColor *string `json:"chartColor,omitempty"`
	ChartDirection *string `json:"chartDirection,omitempty"`
	ChartRankBy *string `json:"chartRankBy,omitempty"`
	ChartType *string `json:"chartType,omitempty"`
	Colspan *int32 `json:"colspan,omitempty"`
	Column *int32 `json:"column,omitempty"`
	HideEmptyRows *bool `json:"hideEmptyRows,omitempty"`
	Id *string `json:"id,omitempty"`
	Metric *UiSettingsTileMetric `json:"metric,omitempty"`
	Name *string `json:"name,omitempty"`
	Row *int32 `json:"row,omitempty"`
	Rowspan *int32 `json:"rowspan,omitempty"`
	ScopeId *string `json:"scopeId,omitempty"`
	ScopeType *string `json:"scopeType,omitempty"`
	SortedColumnIds []string `json:"sortedColumnIds,omitempty"`
	TimeRange *UiSettingsTileTimeRange `json:"timeRange,omitempty"`
	TrendType *string `json:"trendType,omitempty"`
	VizType *string `json:"vizType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiSettingsTile UiSettingsTile

// NewUiSettingsTile instantiates a new UiSettingsTile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiSettingsTile() *UiSettingsTile {
	this := UiSettingsTile{}
	return &this
}

// NewUiSettingsTileWithDefaults instantiates a new UiSettingsTile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiSettingsTileWithDefaults() *UiSettingsTile {
	this := UiSettingsTile{}
	return &this
}

// GetChartBand returns the ChartBand field value if set, zero value otherwise.
func (o *UiSettingsTile) GetChartBand() string {
	if o == nil || IsNil(o.ChartBand) {
		var ret string
		return ret
	}
	return *o.ChartBand
}

// GetChartBandOk returns a tuple with the ChartBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetChartBandOk() (*string, bool) {
	if o == nil || IsNil(o.ChartBand) {
		return nil, false
	}
	return o.ChartBand, true
}

// HasChartBand returns a boolean if a field has been set.
func (o *UiSettingsTile) HasChartBand() bool {
	if o != nil && !IsNil(o.ChartBand) {
		return true
	}

	return false
}

// SetChartBand gets a reference to the given string and assigns it to the ChartBand field.
func (o *UiSettingsTile) SetChartBand(v string) {
	o.ChartBand = &v
}

// GetChartColor returns the ChartColor field value if set, zero value otherwise.
func (o *UiSettingsTile) GetChartColor() string {
	if o == nil || IsNil(o.ChartColor) {
		var ret string
		return ret
	}
	return *o.ChartColor
}

// GetChartColorOk returns a tuple with the ChartColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetChartColorOk() (*string, bool) {
	if o == nil || IsNil(o.ChartColor) {
		return nil, false
	}
	return o.ChartColor, true
}

// HasChartColor returns a boolean if a field has been set.
func (o *UiSettingsTile) HasChartColor() bool {
	if o != nil && !IsNil(o.ChartColor) {
		return true
	}

	return false
}

// SetChartColor gets a reference to the given string and assigns it to the ChartColor field.
func (o *UiSettingsTile) SetChartColor(v string) {
	o.ChartColor = &v
}

// GetChartDirection returns the ChartDirection field value if set, zero value otherwise.
func (o *UiSettingsTile) GetChartDirection() string {
	if o == nil || IsNil(o.ChartDirection) {
		var ret string
		return ret
	}
	return *o.ChartDirection
}

// GetChartDirectionOk returns a tuple with the ChartDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetChartDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.ChartDirection) {
		return nil, false
	}
	return o.ChartDirection, true
}

// HasChartDirection returns a boolean if a field has been set.
func (o *UiSettingsTile) HasChartDirection() bool {
	if o != nil && !IsNil(o.ChartDirection) {
		return true
	}

	return false
}

// SetChartDirection gets a reference to the given string and assigns it to the ChartDirection field.
func (o *UiSettingsTile) SetChartDirection(v string) {
	o.ChartDirection = &v
}

// GetChartRankBy returns the ChartRankBy field value if set, zero value otherwise.
func (o *UiSettingsTile) GetChartRankBy() string {
	if o == nil || IsNil(o.ChartRankBy) {
		var ret string
		return ret
	}
	return *o.ChartRankBy
}

// GetChartRankByOk returns a tuple with the ChartRankBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetChartRankByOk() (*string, bool) {
	if o == nil || IsNil(o.ChartRankBy) {
		return nil, false
	}
	return o.ChartRankBy, true
}

// HasChartRankBy returns a boolean if a field has been set.
func (o *UiSettingsTile) HasChartRankBy() bool {
	if o != nil && !IsNil(o.ChartRankBy) {
		return true
	}

	return false
}

// SetChartRankBy gets a reference to the given string and assigns it to the ChartRankBy field.
func (o *UiSettingsTile) SetChartRankBy(v string) {
	o.ChartRankBy = &v
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *UiSettingsTile) GetChartType() string {
	if o == nil || IsNil(o.ChartType) {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetChartTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChartType) {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *UiSettingsTile) HasChartType() bool {
	if o != nil && !IsNil(o.ChartType) {
		return true
	}

	return false
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *UiSettingsTile) SetChartType(v string) {
	o.ChartType = &v
}

// GetColspan returns the Colspan field value if set, zero value otherwise.
func (o *UiSettingsTile) GetColspan() int32 {
	if o == nil || IsNil(o.Colspan) {
		var ret int32
		return ret
	}
	return *o.Colspan
}

// GetColspanOk returns a tuple with the Colspan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetColspanOk() (*int32, bool) {
	if o == nil || IsNil(o.Colspan) {
		return nil, false
	}
	return o.Colspan, true
}

// HasColspan returns a boolean if a field has been set.
func (o *UiSettingsTile) HasColspan() bool {
	if o != nil && !IsNil(o.Colspan) {
		return true
	}

	return false
}

// SetColspan gets a reference to the given int32 and assigns it to the Colspan field.
func (o *UiSettingsTile) SetColspan(v int32) {
	o.Colspan = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *UiSettingsTile) GetColumn() int32 {
	if o == nil || IsNil(o.Column) {
		var ret int32
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetColumnOk() (*int32, bool) {
	if o == nil || IsNil(o.Column) {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *UiSettingsTile) HasColumn() bool {
	if o != nil && !IsNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int32 and assigns it to the Column field.
func (o *UiSettingsTile) SetColumn(v int32) {
	o.Column = &v
}

// GetHideEmptyRows returns the HideEmptyRows field value if set, zero value otherwise.
func (o *UiSettingsTile) GetHideEmptyRows() bool {
	if o == nil || IsNil(o.HideEmptyRows) {
		var ret bool
		return ret
	}
	return *o.HideEmptyRows
}

// GetHideEmptyRowsOk returns a tuple with the HideEmptyRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetHideEmptyRowsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideEmptyRows) {
		return nil, false
	}
	return o.HideEmptyRows, true
}

// HasHideEmptyRows returns a boolean if a field has been set.
func (o *UiSettingsTile) HasHideEmptyRows() bool {
	if o != nil && !IsNil(o.HideEmptyRows) {
		return true
	}

	return false
}

// SetHideEmptyRows gets a reference to the given bool and assigns it to the HideEmptyRows field.
func (o *UiSettingsTile) SetHideEmptyRows(v bool) {
	o.HideEmptyRows = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UiSettingsTile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UiSettingsTile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UiSettingsTile) SetId(v string) {
	o.Id = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *UiSettingsTile) GetMetric() UiSettingsTileMetric {
	if o == nil || IsNil(o.Metric) {
		var ret UiSettingsTileMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetMetricOk() (*UiSettingsTileMetric, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *UiSettingsTile) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given UiSettingsTileMetric and assigns it to the Metric field.
func (o *UiSettingsTile) SetMetric(v UiSettingsTileMetric) {
	o.Metric = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UiSettingsTile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UiSettingsTile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UiSettingsTile) SetName(v string) {
	o.Name = &v
}

// GetRow returns the Row field value if set, zero value otherwise.
func (o *UiSettingsTile) GetRow() int32 {
	if o == nil || IsNil(o.Row) {
		var ret int32
		return ret
	}
	return *o.Row
}

// GetRowOk returns a tuple with the Row field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetRowOk() (*int32, bool) {
	if o == nil || IsNil(o.Row) {
		return nil, false
	}
	return o.Row, true
}

// HasRow returns a boolean if a field has been set.
func (o *UiSettingsTile) HasRow() bool {
	if o != nil && !IsNil(o.Row) {
		return true
	}

	return false
}

// SetRow gets a reference to the given int32 and assigns it to the Row field.
func (o *UiSettingsTile) SetRow(v int32) {
	o.Row = &v
}

// GetRowspan returns the Rowspan field value if set, zero value otherwise.
func (o *UiSettingsTile) GetRowspan() int32 {
	if o == nil || IsNil(o.Rowspan) {
		var ret int32
		return ret
	}
	return *o.Rowspan
}

// GetRowspanOk returns a tuple with the Rowspan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetRowspanOk() (*int32, bool) {
	if o == nil || IsNil(o.Rowspan) {
		return nil, false
	}
	return o.Rowspan, true
}

// HasRowspan returns a boolean if a field has been set.
func (o *UiSettingsTile) HasRowspan() bool {
	if o != nil && !IsNil(o.Rowspan) {
		return true
	}

	return false
}

// SetRowspan gets a reference to the given int32 and assigns it to the Rowspan field.
func (o *UiSettingsTile) SetRowspan(v int32) {
	o.Rowspan = &v
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise.
func (o *UiSettingsTile) GetScopeId() string {
	if o == nil || IsNil(o.ScopeId) {
		var ret string
		return ret
	}
	return *o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetScopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeId) {
		return nil, false
	}
	return o.ScopeId, true
}

// HasScopeId returns a boolean if a field has been set.
func (o *UiSettingsTile) HasScopeId() bool {
	if o != nil && !IsNil(o.ScopeId) {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given string and assigns it to the ScopeId field.
func (o *UiSettingsTile) SetScopeId(v string) {
	o.ScopeId = &v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *UiSettingsTile) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType) {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetScopeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeType) {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *UiSettingsTile) HasScopeType() bool {
	if o != nil && !IsNil(o.ScopeType) {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *UiSettingsTile) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetSortedColumnIds returns the SortedColumnIds field value if set, zero value otherwise.
func (o *UiSettingsTile) GetSortedColumnIds() []string {
	if o == nil || IsNil(o.SortedColumnIds) {
		var ret []string
		return ret
	}
	return o.SortedColumnIds
}

// GetSortedColumnIdsOk returns a tuple with the SortedColumnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetSortedColumnIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SortedColumnIds) {
		return nil, false
	}
	return o.SortedColumnIds, true
}

// HasSortedColumnIds returns a boolean if a field has been set.
func (o *UiSettingsTile) HasSortedColumnIds() bool {
	if o != nil && !IsNil(o.SortedColumnIds) {
		return true
	}

	return false
}

// SetSortedColumnIds gets a reference to the given []string and assigns it to the SortedColumnIds field.
func (o *UiSettingsTile) SetSortedColumnIds(v []string) {
	o.SortedColumnIds = v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *UiSettingsTile) GetTimeRange() UiSettingsTileTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret UiSettingsTileTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetTimeRangeOk() (*UiSettingsTileTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *UiSettingsTile) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given UiSettingsTileTimeRange and assigns it to the TimeRange field.
func (o *UiSettingsTile) SetTimeRange(v UiSettingsTileTimeRange) {
	o.TimeRange = &v
}

// GetTrendType returns the TrendType field value if set, zero value otherwise.
func (o *UiSettingsTile) GetTrendType() string {
	if o == nil || IsNil(o.TrendType) {
		var ret string
		return ret
	}
	return *o.TrendType
}

// GetTrendTypeOk returns a tuple with the TrendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetTrendTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrendType) {
		return nil, false
	}
	return o.TrendType, true
}

// HasTrendType returns a boolean if a field has been set.
func (o *UiSettingsTile) HasTrendType() bool {
	if o != nil && !IsNil(o.TrendType) {
		return true
	}

	return false
}

// SetTrendType gets a reference to the given string and assigns it to the TrendType field.
func (o *UiSettingsTile) SetTrendType(v string) {
	o.TrendType = &v
}

// GetVizType returns the VizType field value if set, zero value otherwise.
func (o *UiSettingsTile) GetVizType() string {
	if o == nil || IsNil(o.VizType) {
		var ret string
		return ret
	}
	return *o.VizType
}

// GetVizTypeOk returns a tuple with the VizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiSettingsTile) GetVizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VizType) {
		return nil, false
	}
	return o.VizType, true
}

// HasVizType returns a boolean if a field has been set.
func (o *UiSettingsTile) HasVizType() bool {
	if o != nil && !IsNil(o.VizType) {
		return true
	}

	return false
}

// SetVizType gets a reference to the given string and assigns it to the VizType field.
func (o *UiSettingsTile) SetVizType(v string) {
	o.VizType = &v
}

func (o UiSettingsTile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiSettingsTile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartBand) {
		toSerialize["chartBand"] = o.ChartBand
	}
	if !IsNil(o.ChartColor) {
		toSerialize["chartColor"] = o.ChartColor
	}
	if !IsNil(o.ChartDirection) {
		toSerialize["chartDirection"] = o.ChartDirection
	}
	if !IsNil(o.ChartRankBy) {
		toSerialize["chartRankBy"] = o.ChartRankBy
	}
	if !IsNil(o.ChartType) {
		toSerialize["chartType"] = o.ChartType
	}
	if !IsNil(o.Colspan) {
		toSerialize["colspan"] = o.Colspan
	}
	if !IsNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	if !IsNil(o.HideEmptyRows) {
		toSerialize["hideEmptyRows"] = o.HideEmptyRows
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Row) {
		toSerialize["row"] = o.Row
	}
	if !IsNil(o.Rowspan) {
		toSerialize["rowspan"] = o.Rowspan
	}
	if !IsNil(o.ScopeId) {
		toSerialize["scopeId"] = o.ScopeId
	}
	if !IsNil(o.ScopeType) {
		toSerialize["scopeType"] = o.ScopeType
	}
	if !IsNil(o.SortedColumnIds) {
		toSerialize["sortedColumnIds"] = o.SortedColumnIds
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.TrendType) {
		toSerialize["trendType"] = o.TrendType
	}
	if !IsNil(o.VizType) {
		toSerialize["vizType"] = o.VizType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiSettingsTile) UnmarshalJSON(data []byte) (err error) {
	varUiSettingsTile := _UiSettingsTile{}

	err = json.Unmarshal(data, &varUiSettingsTile)

	if err != nil {
		return err
	}

	*o = UiSettingsTile(varUiSettingsTile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chartBand")
		delete(additionalProperties, "chartColor")
		delete(additionalProperties, "chartDirection")
		delete(additionalProperties, "chartRankBy")
		delete(additionalProperties, "chartType")
		delete(additionalProperties, "colspan")
		delete(additionalProperties, "column")
		delete(additionalProperties, "hideEmptyRows")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "name")
		delete(additionalProperties, "row")
		delete(additionalProperties, "rowspan")
		delete(additionalProperties, "scopeId")
		delete(additionalProperties, "scopeType")
		delete(additionalProperties, "sortedColumnIds")
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "trendType")
		delete(additionalProperties, "vizType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiSettingsTile struct {
	value *UiSettingsTile
	isSet bool
}

func (v NullableUiSettingsTile) Get() *UiSettingsTile {
	return v.value
}

func (v *NullableUiSettingsTile) Set(val *UiSettingsTile) {
	v.value = val
	v.isSet = true
}

func (v NullableUiSettingsTile) IsSet() bool {
	return v.isSet
}

func (v *NullableUiSettingsTile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiSettingsTile(val *UiSettingsTile) *NullableUiSettingsTile {
	return &NullableUiSettingsTile{value: val, isSet: true}
}

func (v NullableUiSettingsTile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiSettingsTile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

