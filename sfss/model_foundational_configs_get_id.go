/*
SmartFabric Storage Software Application REST APIs

REST APIs used for managing SFSS application are captured in this section.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sfss

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FoundationalConfigsGETID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FoundationalConfigsGETID{}

// FoundationalConfigsGETID This model lists the foundational configuration based on the specified CDC instance ID.
type FoundationalConfigsGETID struct {
	// Port on which the CDC listens for mDNS queries; default is 8009
	DiscoveryControllerPort string `json:"DiscoveryControllerPort"`
	// CDC instance identifier
	InstanceIdentifier string `json:"InstanceIdentifier"`
	// A unique NVMe Qualified Name (NQN) that is used to identify the CDC instance.
	NQN string `json:"NQN"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
}

type _FoundationalConfigsGETID FoundationalConfigsGETID

// NewFoundationalConfigsGETID instantiates a new FoundationalConfigsGETID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoundationalConfigsGETID(discoveryControllerPort string, instanceIdentifier string, nQN string, odataId string, odataType string, odataContext string) *FoundationalConfigsGETID {
	this := FoundationalConfigsGETID{}
	this.DiscoveryControllerPort = discoveryControllerPort
	this.InstanceIdentifier = instanceIdentifier
	this.NQN = nQN
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	return &this
}

// NewFoundationalConfigsGETIDWithDefaults instantiates a new FoundationalConfigsGETID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoundationalConfigsGETIDWithDefaults() *FoundationalConfigsGETID {
	this := FoundationalConfigsGETID{}
	return &this
}

// GetDiscoveryControllerPort returns the DiscoveryControllerPort field value
func (o *FoundationalConfigsGETID) GetDiscoveryControllerPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscoveryControllerPort
}

// GetDiscoveryControllerPortOk returns a tuple with the DiscoveryControllerPort field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetDiscoveryControllerPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryControllerPort, true
}

// SetDiscoveryControllerPort sets field value
func (o *FoundationalConfigsGETID) SetDiscoveryControllerPort(v string) {
	o.DiscoveryControllerPort = v
}

// GetInstanceIdentifier returns the InstanceIdentifier field value
func (o *FoundationalConfigsGETID) GetInstanceIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceIdentifier
}

// GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetInstanceIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceIdentifier, true
}

// SetInstanceIdentifier sets field value
func (o *FoundationalConfigsGETID) SetInstanceIdentifier(v string) {
	o.InstanceIdentifier = v
}

// GetNQN returns the NQN field value
func (o *FoundationalConfigsGETID) GetNQN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NQN
}

// GetNQNOk returns a tuple with the NQN field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetNQNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NQN, true
}

// SetNQN sets field value
func (o *FoundationalConfigsGETID) SetNQN(v string) {
	o.NQN = v
}

// GetOdataId returns the OdataId field value
func (o *FoundationalConfigsGETID) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *FoundationalConfigsGETID) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *FoundationalConfigsGETID) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *FoundationalConfigsGETID) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *FoundationalConfigsGETID) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGETID) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *FoundationalConfigsGETID) SetOdataContext(v string) {
	o.OdataContext = v
}

func (o FoundationalConfigsGETID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FoundationalConfigsGETID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DiscoveryControllerPort"] = o.DiscoveryControllerPort
	toSerialize["InstanceIdentifier"] = o.InstanceIdentifier
	toSerialize["NQN"] = o.NQN
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	return toSerialize, nil
}

func (o *FoundationalConfigsGETID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DiscoveryControllerPort",
		"InstanceIdentifier",
		"NQN",
		"@odata.id",
		"@odata.type",
		"@odata.context",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFoundationalConfigsGETID := _FoundationalConfigsGETID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFoundationalConfigsGETID)

	if err != nil {
		return err
	}

	*o = FoundationalConfigsGETID(varFoundationalConfigsGETID)

	return err
}

type NullableFoundationalConfigsGETID struct {
	value *FoundationalConfigsGETID
	isSet bool
}

func (v NullableFoundationalConfigsGETID) Get() *FoundationalConfigsGETID {
	return v.value
}

func (v *NullableFoundationalConfigsGETID) Set(val *FoundationalConfigsGETID) {
	v.value = val
	v.isSet = true
}

func (v NullableFoundationalConfigsGETID) IsSet() bool {
	return v.isSet
}

func (v *NullableFoundationalConfigsGETID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoundationalConfigsGETID(val *FoundationalConfigsGETID) *NullableFoundationalConfigsGETID {
	return &NullableFoundationalConfigsGETID{value: val, isSet: true}
}

func (v NullableFoundationalConfigsGETID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoundationalConfigsGETID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


