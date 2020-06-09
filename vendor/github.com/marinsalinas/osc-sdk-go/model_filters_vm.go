/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// FiltersVm One or more filters.
type FiltersVm struct {
	// The keys of the tags associated with the VMs.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the VMs.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the VMs, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
	// One or more IDs of VMs.
	VmIds *[]string `json:"VmIds,omitempty"`
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersVm) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVm) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersVm) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersVm) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersVm) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVm) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersVm) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersVm) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersVm) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVm) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersVm) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersVm) SetTags(v []string) {
	o.Tags = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *FiltersVm) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersVm) GetVmIdsOk() ([]string, bool) {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret, false
	}
	return *o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *FiltersVm) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *FiltersVm) SetVmIds(v []string) {
	o.VmIds = &v
}

type NullableFiltersVm struct {
	Value        FiltersVm
	ExplicitNull bool
}

func (v NullableFiltersVm) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersVm) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
