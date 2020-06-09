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

// CreateImageResponse struct for CreateImageResponse
type CreateImageResponse struct {
	Image           *Image           `json:"Image,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateImageResponse) GetImage() Image {
	if o == nil || o.Image == nil {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageResponse) GetImageOk() (Image, bool) {
	if o == nil || o.Image == nil {
		var ret Image
		return ret, false
	}
	return *o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateImageResponse) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *CreateImageResponse) SetImage(v Image) {
	o.Image = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateImageResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateImageResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateImageResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateImageResponse struct {
	Value        CreateImageResponse
	ExplicitNull bool
}

func (v NullableCreateImageResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateImageResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
