/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the SloResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloResponse{}

// SloResponse struct for SloResponse
type SloResponse struct {
	// The identifier of the SLO.
	Id *string `json:"id,omitempty"`
	// The name of the SLO.
	Name *string `json:"name,omitempty"`
	// The description of the SLO.
	Description     *string                `json:"description,omitempty"`
	Indicator       *SloResponseIndicator  `json:"indicator,omitempty"`
	TimeWindow      *SloResponseTimeWindow `json:"timeWindow,omitempty"`
	BudgetingMethod *BudgetingMethod       `json:"budgetingMethod,omitempty"`
	Objective       *Objective             `json:"objective,omitempty"`
	Settings        *Settings              `json:"settings,omitempty"`
	// The SLO revision
	Revision *float32 `json:"revision,omitempty"`
	Summary  *Summary `json:"summary,omitempty"`
	// Indicate if the SLO is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The creation date
	CreatedAt *string `json:"createdAt,omitempty"`
	// The last update date
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewSloResponse instantiates a new SloResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloResponse() *SloResponse {
	this := SloResponse{}
	return &this
}

// NewSloResponseWithDefaults instantiates a new SloResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloResponseWithDefaults() *SloResponse {
	this := SloResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SloResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SloResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SloResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SloResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SloResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SloResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SloResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SloResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SloResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *SloResponse) GetIndicator() SloResponseIndicator {
	if o == nil || IsNil(o.Indicator) {
		var ret SloResponseIndicator
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetIndicatorOk() (*SloResponseIndicator, bool) {
	if o == nil || IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *SloResponse) HasIndicator() bool {
	if o != nil && !IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given SloResponseIndicator and assigns it to the Indicator field.
func (o *SloResponse) SetIndicator(v SloResponseIndicator) {
	o.Indicator = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *SloResponse) GetTimeWindow() SloResponseTimeWindow {
	if o == nil || IsNil(o.TimeWindow) {
		var ret SloResponseTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetTimeWindowOk() (*SloResponseTimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *SloResponse) HasTimeWindow() bool {
	if o != nil && !IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given SloResponseTimeWindow and assigns it to the TimeWindow field.
func (o *SloResponse) SetTimeWindow(v SloResponseTimeWindow) {
	o.TimeWindow = &v
}

// GetBudgetingMethod returns the BudgetingMethod field value if set, zero value otherwise.
func (o *SloResponse) GetBudgetingMethod() BudgetingMethod {
	if o == nil || IsNil(o.BudgetingMethod) {
		var ret BudgetingMethod
		return ret
	}
	return *o.BudgetingMethod
}

// GetBudgetingMethodOk returns a tuple with the BudgetingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetBudgetingMethodOk() (*BudgetingMethod, bool) {
	if o == nil || IsNil(o.BudgetingMethod) {
		return nil, false
	}
	return o.BudgetingMethod, true
}

// HasBudgetingMethod returns a boolean if a field has been set.
func (o *SloResponse) HasBudgetingMethod() bool {
	if o != nil && !IsNil(o.BudgetingMethod) {
		return true
	}

	return false
}

// SetBudgetingMethod gets a reference to the given BudgetingMethod and assigns it to the BudgetingMethod field.
func (o *SloResponse) SetBudgetingMethod(v BudgetingMethod) {
	o.BudgetingMethod = &v
}

// GetObjective returns the Objective field value if set, zero value otherwise.
func (o *SloResponse) GetObjective() Objective {
	if o == nil || IsNil(o.Objective) {
		var ret Objective
		return ret
	}
	return *o.Objective
}

// GetObjectiveOk returns a tuple with the Objective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetObjectiveOk() (*Objective, bool) {
	if o == nil || IsNil(o.Objective) {
		return nil, false
	}
	return o.Objective, true
}

// HasObjective returns a boolean if a field has been set.
func (o *SloResponse) HasObjective() bool {
	if o != nil && !IsNil(o.Objective) {
		return true
	}

	return false
}

// SetObjective gets a reference to the given Objective and assigns it to the Objective field.
func (o *SloResponse) SetObjective(v Objective) {
	o.Objective = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SloResponse) GetSettings() Settings {
	if o == nil || IsNil(o.Settings) {
		var ret Settings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetSettingsOk() (*Settings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SloResponse) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Settings and assigns it to the Settings field.
func (o *SloResponse) SetSettings(v Settings) {
	o.Settings = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *SloResponse) GetRevision() float32 {
	if o == nil || IsNil(o.Revision) {
		var ret float32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetRevisionOk() (*float32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *SloResponse) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given float32 and assigns it to the Revision field.
func (o *SloResponse) SetRevision(v float32) {
	o.Revision = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *SloResponse) GetSummary() Summary {
	if o == nil || IsNil(o.Summary) {
		var ret Summary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetSummaryOk() (*Summary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *SloResponse) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given Summary and assigns it to the Summary field.
func (o *SloResponse) SetSummary(v Summary) {
	o.Summary = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SloResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SloResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SloResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SloResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SloResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SloResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SloResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SloResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SloResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o SloResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Indicator) {
		toSerialize["indicator"] = o.Indicator
	}
	if !IsNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	if !IsNil(o.BudgetingMethod) {
		toSerialize["budgetingMethod"] = o.BudgetingMethod
	}
	if !IsNil(o.Objective) {
		toSerialize["objective"] = o.Objective
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSloResponse struct {
	value *SloResponse
	isSet bool
}

func (v NullableSloResponse) Get() *SloResponse {
	return v.value
}

func (v *NullableSloResponse) Set(val *SloResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSloResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSloResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloResponse(val *SloResponse) *NullableSloResponse {
	return &NullableSloResponse{value: val, isSet: true}
}

func (v NullableSloResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}