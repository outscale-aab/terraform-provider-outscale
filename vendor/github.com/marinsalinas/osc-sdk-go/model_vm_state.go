/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// VmState Information about the state of the VM.
type VmState struct {
	// The current state of the VM (`InService` \\| `OutOfService` \\| `Unknown`).
	CurrentState *string `json:"CurrentState,omitempty"`
	// The previous state of the VM (`InService` \\| `OutOfService` \\| `Unknown`).
	PreviousState *string `json:"PreviousState,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *VmState) GetCurrentState() string {
	if o == nil || o.CurrentState == nil {
		var ret string
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetCurrentStateOk() (string, bool) {
	if o == nil || o.CurrentState == nil {
		var ret string
		return ret, false
	}
	return *o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *VmState) HasCurrentState() bool {
	if o != nil && o.CurrentState != nil {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.
func (o *VmState) SetCurrentState(v string) {
	o.CurrentState = &v
}

// GetPreviousState returns the PreviousState field value if set, zero value otherwise.
func (o *VmState) GetPreviousState() string {
	if o == nil || o.PreviousState == nil {
		var ret string
		return ret
	}
	return *o.PreviousState
}

// GetPreviousStateOk returns a tuple with the PreviousState field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetPreviousStateOk() (string, bool) {
	if o == nil || o.PreviousState == nil {
		var ret string
		return ret, false
	}
	return *o.PreviousState, true
}

// HasPreviousState returns a boolean if a field has been set.
func (o *VmState) HasPreviousState() bool {
	if o != nil && o.PreviousState != nil {
		return true
	}

	return false
}

// SetPreviousState gets a reference to the given string and assigns it to the PreviousState field.
func (o *VmState) SetPreviousState(v string) {
	o.PreviousState = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VmState) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *VmState) GetVmIdOk() (string, bool) {
	if o == nil || o.VmId == nil {
		var ret string
		return ret, false
	}
	return *o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VmState) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *VmState) SetVmId(v string) {
	o.VmId = &v
}

type NullableVmState struct {
	Value        VmState
	ExplicitNull bool
}

func (v NullableVmState) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableVmState) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}