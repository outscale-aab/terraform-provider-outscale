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

// CreateNatServiceResponse struct for CreateNatServiceResponse
type CreateNatServiceResponse struct {
	NatService      *NatService      `json:"NatService,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNatService returns the NatService field value if set, zero value otherwise.
func (o *CreateNatServiceResponse) GetNatService() NatService {
	if o == nil || o.NatService == nil {
		var ret NatService
		return ret
	}
	return *o.NatService
}

// GetNatServiceOk returns a tuple with the NatService field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceResponse) GetNatServiceOk() (NatService, bool) {
	if o == nil || o.NatService == nil {
		var ret NatService
		return ret, false
	}
	return *o.NatService, true
}

// HasNatService returns a boolean if a field has been set.
func (o *CreateNatServiceResponse) HasNatService() bool {
	if o != nil && o.NatService != nil {
		return true
	}

	return false
}

// SetNatService gets a reference to the given NatService and assigns it to the NatService field.
func (o *CreateNatServiceResponse) SetNatService(v NatService) {
	o.NatService = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateNatServiceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateNatServiceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateNatServiceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateNatServiceResponse struct {
	Value        CreateNatServiceResponse
	ExplicitNull bool
}

func (v NullableCreateNatServiceResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateNatServiceResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
