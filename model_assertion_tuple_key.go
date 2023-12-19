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

// AssertionTupleKey struct for AssertionTupleKey
type AssertionTupleKey struct {
	Object   string `json:"object"yaml:"object"`
	Relation string `json:"relation"yaml:"relation"`
	User     string `json:"user"yaml:"user"`
}

// NewAssertionTupleKey instantiates a new AssertionTupleKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertionTupleKey(object string, relation string, user string) *AssertionTupleKey {
	this := AssertionTupleKey{}
	this.Object = object
	this.Relation = relation
	this.User = user
	return &this
}

// NewAssertionTupleKeyWithDefaults instantiates a new AssertionTupleKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertionTupleKeyWithDefaults() *AssertionTupleKey {
	this := AssertionTupleKey{}
	return &this
}

// GetObject returns the Object field value
func (o *AssertionTupleKey) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AssertionTupleKey) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AssertionTupleKey) SetObject(v string) {
	o.Object = v
}

// GetRelation returns the Relation field value
func (o *AssertionTupleKey) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *AssertionTupleKey) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *AssertionTupleKey) SetRelation(v string) {
	o.Relation = v
}

// GetUser returns the User field value
func (o *AssertionTupleKey) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AssertionTupleKey) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AssertionTupleKey) SetUser(v string) {
	o.User = v
}

func (o AssertionTupleKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["relation"] = o.Relation
	toSerialize["user"] = o.User
	return json.Marshal(toSerialize)
}

type NullableAssertionTupleKey struct {
	value *AssertionTupleKey
	isSet bool
}

func (v NullableAssertionTupleKey) Get() *AssertionTupleKey {
	return v.value
}

func (v *NullableAssertionTupleKey) Set(val *AssertionTupleKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertionTupleKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertionTupleKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertionTupleKey(val *AssertionTupleKey) *NullableAssertionTupleKey {
	return &NullableAssertionTupleKey{value: val, isSet: true}
}

func (v NullableAssertionTupleKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertionTupleKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}