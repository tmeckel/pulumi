/*
Pulumi Service API

The Pulumi Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Stack Stack describes a Stack running on a Pulumi Cloud.
type Stack struct {
	OrgName *string `json:"orgName,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
	StackName *string `json:"stackName,omitempty"`
	ActiveUpdate *string `json:"activeUpdate,omitempty"`
	// Tags are a set of key values applied to stacks.
	Tags *map[string]string `json:"tags,omitempty"`
	Version *int32 `json:"version,omitempty"`
	CurrentOperation *OperationStatus `json:"currentOperation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Stack Stack

// NewStack instantiates a new Stack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStack() *Stack {
	this := Stack{}
	return &this
}

// NewStackWithDefaults instantiates a new Stack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackWithDefaults() *Stack {
	this := Stack{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Stack) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Stack) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Stack) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *Stack) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *Stack) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *Stack) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetStackName returns the StackName field value if set, zero value otherwise.
func (o *Stack) GetStackName() string {
	if o == nil || o.StackName == nil {
		var ret string
		return ret
	}
	return *o.StackName
}

// GetStackNameOk returns a tuple with the StackName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetStackNameOk() (*string, bool) {
	if o == nil || o.StackName == nil {
		return nil, false
	}
	return o.StackName, true
}

// HasStackName returns a boolean if a field has been set.
func (o *Stack) HasStackName() bool {
	if o != nil && o.StackName != nil {
		return true
	}

	return false
}

// SetStackName gets a reference to the given string and assigns it to the StackName field.
func (o *Stack) SetStackName(v string) {
	o.StackName = &v
}

// GetActiveUpdate returns the ActiveUpdate field value if set, zero value otherwise.
func (o *Stack) GetActiveUpdate() string {
	if o == nil || o.ActiveUpdate == nil {
		var ret string
		return ret
	}
	return *o.ActiveUpdate
}

// GetActiveUpdateOk returns a tuple with the ActiveUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetActiveUpdateOk() (*string, bool) {
	if o == nil || o.ActiveUpdate == nil {
		return nil, false
	}
	return o.ActiveUpdate, true
}

// HasActiveUpdate returns a boolean if a field has been set.
func (o *Stack) HasActiveUpdate() bool {
	if o != nil && o.ActiveUpdate != nil {
		return true
	}

	return false
}

// SetActiveUpdate gets a reference to the given string and assigns it to the ActiveUpdate field.
func (o *Stack) SetActiveUpdate(v string) {
	o.ActiveUpdate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Stack) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Stack) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Stack) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Stack) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Stack) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Stack) SetVersion(v int32) {
	o.Version = &v
}

// GetCurrentOperation returns the CurrentOperation field value if set, zero value otherwise.
func (o *Stack) GetCurrentOperation() OperationStatus {
	if o == nil || o.CurrentOperation == nil {
		var ret OperationStatus
		return ret
	}
	return *o.CurrentOperation
}

// GetCurrentOperationOk returns a tuple with the CurrentOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stack) GetCurrentOperationOk() (*OperationStatus, bool) {
	if o == nil || o.CurrentOperation == nil {
		return nil, false
	}
	return o.CurrentOperation, true
}

// HasCurrentOperation returns a boolean if a field has been set.
func (o *Stack) HasCurrentOperation() bool {
	if o != nil && o.CurrentOperation != nil {
		return true
	}

	return false
}

// SetCurrentOperation gets a reference to the given OperationStatus and assigns it to the CurrentOperation field.
func (o *Stack) SetCurrentOperation(v OperationStatus) {
	o.CurrentOperation = &v
}

func (o Stack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ProjectName != nil {
		toSerialize["projectName"] = o.ProjectName
	}
	if o.StackName != nil {
		toSerialize["stackName"] = o.StackName
	}
	if o.ActiveUpdate != nil {
		toSerialize["activeUpdate"] = o.ActiveUpdate
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.CurrentOperation != nil {
		toSerialize["currentOperation"] = o.CurrentOperation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Stack) UnmarshalJSON(bytes []byte) (err error) {
	varStack := _Stack{}

	if err = json.Unmarshal(bytes, &varStack); err == nil {
		*o = Stack(varStack)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "projectName")
		delete(additionalProperties, "stackName")
		delete(additionalProperties, "activeUpdate")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "version")
		delete(additionalProperties, "currentOperation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStack struct {
	value *Stack
	isSet bool
}

func (v NullableStack) Get() *Stack {
	return v.value
}

func (v *NullableStack) Set(val *Stack) {
	v.value = val
	v.isSet = true
}

func (v NullableStack) IsSet() bool {
	return v.isSet
}

func (v *NullableStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStack(val *Stack) *NullableStack {
	return &NullableStack{value: val, isSet: true}
}

func (v NullableStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


