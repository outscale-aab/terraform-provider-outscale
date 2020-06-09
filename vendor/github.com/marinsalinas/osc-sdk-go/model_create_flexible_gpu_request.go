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

// CreateFlexibleGpuRequest struct for CreateFlexibleGpuRequest
type CreateFlexibleGpuRequest struct {
	// If `true`, the fGPU is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by [ReadFlexibleGpuCatalog](#readflexiblegpucatalog) for the specified model of fGPU).
	Generation *string `json:"Generation,omitempty"`
	// The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).
	ModelName string `json:"ModelName"`
	// The Subregion in which you want to create the fGPU.
	SubregionName string `json:"SubregionName"`
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret, false
	}
	return *o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *CreateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *CreateFlexibleGpuRequest) GetGeneration() string {
	if o == nil || o.Generation == nil {
		var ret string
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuRequest) GetGenerationOk() (string, bool) {
	if o == nil || o.Generation == nil {
		var ret string
		return ret, false
	}
	return *o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *CreateFlexibleGpuRequest) HasGeneration() bool {
	if o != nil && o.Generation != nil {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given string and assigns it to the Generation field.
func (o *CreateFlexibleGpuRequest) SetGeneration(v string) {
	o.Generation = &v
}

// GetModelName returns the ModelName field value
func (o *CreateFlexibleGpuRequest) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// SetModelName sets field value
func (o *CreateFlexibleGpuRequest) SetModelName(v string) {
	o.ModelName = v
}

// GetSubregionName returns the SubregionName field value
func (o *CreateFlexibleGpuRequest) GetSubregionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubregionName
}

// SetSubregionName sets field value
func (o *CreateFlexibleGpuRequest) SetSubregionName(v string) {
	o.SubregionName = v
}

type NullableCreateFlexibleGpuRequest struct {
	Value        CreateFlexibleGpuRequest
	ExplicitNull bool
}

func (v NullableCreateFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
