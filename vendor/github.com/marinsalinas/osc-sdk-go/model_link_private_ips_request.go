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

// LinkPrivateIpsRequest struct for LinkPrivateIpsRequest
type LinkPrivateIpsRequest struct {
	// If `true`, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.
	AllowRelink *bool `json:"AllowRelink,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId"`
	// The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet.
	PrivateIps *[]string `json:"PrivateIps,omitempty"`
	// The number of secondary private IP addresses to assign to the NIC.
	SecondaryPrivateIpCount *int64 `json:"SecondaryPrivateIpCount,omitempty"`
}

// GetAllowRelink returns the AllowRelink field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetAllowRelink() bool {
	if o == nil || o.AllowRelink == nil {
		var ret bool
		return ret
	}
	return *o.AllowRelink
}

// GetAllowRelinkOk returns a tuple with the AllowRelink field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetAllowRelinkOk() (bool, bool) {
	if o == nil || o.AllowRelink == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowRelink, true
}

// HasAllowRelink returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasAllowRelink() bool {
	if o != nil && o.AllowRelink != nil {
		return true
	}

	return false
}

// SetAllowRelink gets a reference to the given bool and assigns it to the AllowRelink field.
func (o *LinkPrivateIpsRequest) SetAllowRelink(v bool) {
	o.AllowRelink = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkPrivateIpsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value
func (o *LinkPrivateIpsRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// SetNicId sets field value
func (o *LinkPrivateIpsRequest) SetNicId(v string) {
	o.NicId = v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetPrivateIps() []string {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetPrivateIpsOk() ([]string, bool) {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret, false
	}
	return *o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.
func (o *LinkPrivateIpsRequest) SetPrivateIps(v []string) {
	o.PrivateIps = &v
}

// GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field value if set, zero value otherwise.
func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCount() int64 {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int64
		return ret
	}
	return *o.SecondaryPrivateIpCount
}

// GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCountOk() (int64, bool) {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int64
		return ret, false
	}
	return *o.SecondaryPrivateIpCount, true
}

// HasSecondaryPrivateIpCount returns a boolean if a field has been set.
func (o *LinkPrivateIpsRequest) HasSecondaryPrivateIpCount() bool {
	if o != nil && o.SecondaryPrivateIpCount != nil {
		return true
	}

	return false
}

// SetSecondaryPrivateIpCount gets a reference to the given int64 and assigns it to the SecondaryPrivateIpCount field.
func (o *LinkPrivateIpsRequest) SetSecondaryPrivateIpCount(v int64) {
	o.SecondaryPrivateIpCount = &v
}

type NullableLinkPrivateIpsRequest struct {
	Value        LinkPrivateIpsRequest
	ExplicitNull bool
}

func (v NullableLinkPrivateIpsRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLinkPrivateIpsRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
