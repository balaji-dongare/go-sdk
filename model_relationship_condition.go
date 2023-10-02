/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// RelationshipCondition struct for RelationshipCondition
type RelationshipCondition struct {
	// A reference (by name) of the relationship condition defined in the authorization model.
	ConditionName string `json:"condition_name"`
	// Additional context/data to persist along with the condition. The keys must match the parameters defined by the condition, and the value types must match the parameter type definitions.
	Context *map[string]interface{} `json:"context,omitempty"`
}

// NewRelationshipCondition instantiates a new RelationshipCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipCondition(conditionName string) *RelationshipCondition {
	this := RelationshipCondition{}
	this.ConditionName = conditionName
	return &this
}

// NewRelationshipConditionWithDefaults instantiates a new RelationshipCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipConditionWithDefaults() *RelationshipCondition {
	this := RelationshipCondition{}
	return &this
}

// GetConditionName returns the ConditionName field value
func (o *RelationshipCondition) GetConditionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionName
}

// GetConditionNameOk returns a tuple with the ConditionName field value
// and a boolean to check if the value has been set.
func (o *RelationshipCondition) GetConditionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionName, true
}

// SetConditionName sets field value
func (o *RelationshipCondition) SetConditionName(v string) {
	o.ConditionName = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RelationshipCondition) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipCondition) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RelationshipCondition) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *RelationshipCondition) SetContext(v map[string]interface{}) {
	o.Context = &v
}

func (o RelationshipCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition_name"] = o.ConditionName
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipCondition struct {
	value *RelationshipCondition
	isSet bool
}

func (v NullableRelationshipCondition) Get() *RelationshipCondition {
	return v.value
}

func (v *NullableRelationshipCondition) Set(val *RelationshipCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipCondition(val *RelationshipCondition) *NullableRelationshipCondition {
	return &NullableRelationshipCondition{value: val, isSet: true}
}

func (v NullableRelationshipCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
