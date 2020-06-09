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

// CreateSecurityGroupRequest struct for CreateSecurityGroupRequest
type CreateSecurityGroupRequest struct {
	// A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
	Description string `json:"Description"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net for the security group.
	NetId *string `json:"NetId,omitempty"`
	// (Public Cloud only) The name of the security group.<br /> This name must not start with `sg-`.</br> This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed.
	SecurityGroupName string `json:"SecurityGroupName"`
}

// GetDescription returns the Description field value
func (o *CreateSecurityGroupRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *CreateSecurityGroupRequest) SetDescription(v string) {
	o.Description = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateSecurityGroupRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateSecurityGroupRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateSecurityGroupRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *CreateSecurityGroupRequest) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetNetIdOk() (string, bool) {
	if o == nil || o.NetId == nil {
		var ret string
		return ret, false
	}
	return *o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *CreateSecurityGroupRequest) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *CreateSecurityGroupRequest) SetNetId(v string) {
	o.NetId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value
func (o *CreateSecurityGroupRequest) GetSecurityGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityGroupName
}

// SetSecurityGroupName sets field value
func (o *CreateSecurityGroupRequest) SetSecurityGroupName(v string) {
	o.SecurityGroupName = v
}

type NullableCreateSecurityGroupRequest struct {
	Value        CreateSecurityGroupRequest
	ExplicitNull bool
}

func (v NullableCreateSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateSecurityGroupRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
