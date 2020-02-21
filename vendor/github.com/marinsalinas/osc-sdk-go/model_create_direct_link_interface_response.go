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

// CreateDirectLinkInterfaceResponse struct for CreateDirectLinkInterfaceResponse
type CreateDirectLinkInterfaceResponse struct {
	DirectLinkInterface *DirectLinkInterfaces `json:"DirectLinkInterface,omitempty"`
	ResponseContext     *ResponseContext      `json:"ResponseContext,omitempty"`
}

// GetDirectLinkInterface returns the DirectLinkInterface field value if set, zero value otherwise.
func (o *CreateDirectLinkInterfaceResponse) GetDirectLinkInterface() DirectLinkInterfaces {
	if o == nil || o.DirectLinkInterface == nil {
		var ret DirectLinkInterfaces
		return ret
	}
	return *o.DirectLinkInterface
}

// GetDirectLinkInterfaceOk returns a tuple with the DirectLinkInterface field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceResponse) GetDirectLinkInterfaceOk() (DirectLinkInterfaces, bool) {
	if o == nil || o.DirectLinkInterface == nil {
		var ret DirectLinkInterfaces
		return ret, false
	}
	return *o.DirectLinkInterface, true
}

// HasDirectLinkInterface returns a boolean if a field has been set.
func (o *CreateDirectLinkInterfaceResponse) HasDirectLinkInterface() bool {
	if o != nil && o.DirectLinkInterface != nil {
		return true
	}

	return false
}

// SetDirectLinkInterface gets a reference to the given DirectLinkInterfaces and assigns it to the DirectLinkInterface field.
func (o *CreateDirectLinkInterfaceResponse) SetDirectLinkInterface(v DirectLinkInterfaces) {
	o.DirectLinkInterface = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateDirectLinkInterfaceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateDirectLinkInterfaceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateDirectLinkInterfaceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateDirectLinkInterfaceResponse struct {
	Value        CreateDirectLinkInterfaceResponse
	ExplicitNull bool
}

func (v NullableCreateDirectLinkInterfaceResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateDirectLinkInterfaceResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}