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

// FiltersNic One or more filters.
type FiltersNic struct {
	// The device numbers the NICs are attached to.
	LinkNicSortNumbers *[]int64 `json:"LinkNicSortNumbers,omitempty"`
	// The IDs of the VMs the NICs are attached to.
	LinkNicVmIds *[]string `json:"LinkNicVmIds,omitempty"`
	// The IDs of the NICs.
	NicIds *[]string `json:"NicIds,omitempty"`
	// The private IP addresses of the NICs.
	PrivateIpsPrivateIps *[]string `json:"PrivateIpsPrivateIps,omitempty"`
	// The IDs of the Subnets for the NICs.
	SubnetIds *[]string `json:"SubnetIds,omitempty"`
}

// GetLinkNicSortNumbers returns the LinkNicSortNumbers field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicSortNumbers() []int64 {
	if o == nil || o.LinkNicSortNumbers == nil {
		var ret []int64
		return ret
	}
	return *o.LinkNicSortNumbers
}

// GetLinkNicSortNumbersOk returns a tuple with the LinkNicSortNumbers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicSortNumbersOk() ([]int64, bool) {
	if o == nil || o.LinkNicSortNumbers == nil {
		var ret []int64
		return ret, false
	}
	return *o.LinkNicSortNumbers, true
}

// HasLinkNicSortNumbers returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicSortNumbers() bool {
	if o != nil && o.LinkNicSortNumbers != nil {
		return true
	}

	return false
}

// SetLinkNicSortNumbers gets a reference to the given []int64 and assigns it to the LinkNicSortNumbers field.
func (o *FiltersNic) SetLinkNicSortNumbers(v []int64) {
	o.LinkNicSortNumbers = &v
}

// GetLinkNicVmIds returns the LinkNicVmIds field value if set, zero value otherwise.
func (o *FiltersNic) GetLinkNicVmIds() []string {
	if o == nil || o.LinkNicVmIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkNicVmIds
}

// GetLinkNicVmIdsOk returns a tuple with the LinkNicVmIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetLinkNicVmIdsOk() ([]string, bool) {
	if o == nil || o.LinkNicVmIds == nil {
		var ret []string
		return ret, false
	}
	return *o.LinkNicVmIds, true
}

// HasLinkNicVmIds returns a boolean if a field has been set.
func (o *FiltersNic) HasLinkNicVmIds() bool {
	if o != nil && o.LinkNicVmIds != nil {
		return true
	}

	return false
}

// SetLinkNicVmIds gets a reference to the given []string and assigns it to the LinkNicVmIds field.
func (o *FiltersNic) SetLinkNicVmIds(v []string) {
	o.LinkNicVmIds = &v
}

// GetNicIds returns the NicIds field value if set, zero value otherwise.
func (o *FiltersNic) GetNicIds() []string {
	if o == nil || o.NicIds == nil {
		var ret []string
		return ret
	}
	return *o.NicIds
}

// GetNicIdsOk returns a tuple with the NicIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetNicIdsOk() ([]string, bool) {
	if o == nil || o.NicIds == nil {
		var ret []string
		return ret, false
	}
	return *o.NicIds, true
}

// HasNicIds returns a boolean if a field has been set.
func (o *FiltersNic) HasNicIds() bool {
	if o != nil && o.NicIds != nil {
		return true
	}

	return false
}

// SetNicIds gets a reference to the given []string and assigns it to the NicIds field.
func (o *FiltersNic) SetNicIds(v []string) {
	o.NicIds = &v
}

// GetPrivateIpsPrivateIps returns the PrivateIpsPrivateIps field value if set, zero value otherwise.
func (o *FiltersNic) GetPrivateIpsPrivateIps() []string {
	if o == nil || o.PrivateIpsPrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIpsPrivateIps
}

// GetPrivateIpsPrivateIpsOk returns a tuple with the PrivateIpsPrivateIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetPrivateIpsPrivateIpsOk() ([]string, bool) {
	if o == nil || o.PrivateIpsPrivateIps == nil {
		var ret []string
		return ret, false
	}
	return *o.PrivateIpsPrivateIps, true
}

// HasPrivateIpsPrivateIps returns a boolean if a field has been set.
func (o *FiltersNic) HasPrivateIpsPrivateIps() bool {
	if o != nil && o.PrivateIpsPrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIpsPrivateIps gets a reference to the given []string and assigns it to the PrivateIpsPrivateIps field.
func (o *FiltersNic) SetPrivateIpsPrivateIps(v []string) {
	o.PrivateIpsPrivateIps = &v
}

// GetSubnetIds returns the SubnetIds field value if set, zero value otherwise.
func (o *FiltersNic) GetSubnetIds() []string {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNic) GetSubnetIdsOk() ([]string, bool) {
	if o == nil || o.SubnetIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SubnetIds, true
}

// HasSubnetIds returns a boolean if a field has been set.
func (o *FiltersNic) HasSubnetIds() bool {
	if o != nil && o.SubnetIds != nil {
		return true
	}

	return false
}

// SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.
func (o *FiltersNic) SetSubnetIds(v []string) {
	o.SubnetIds = &v
}

type NullableFiltersNic struct {
	Value        FiltersNic
	ExplicitNull bool
}

func (v NullableFiltersNic) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersNic) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
